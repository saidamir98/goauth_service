package v1

import (
	"time"

	"github.com/saidamir98/goauth_service/config"
	"github.com/saidamir98/goauth_service/pkg/security"
	"github.com/saidamir98/goauth_service/pkg/util"

	"github.com/google/uuid"
	"github.com/saidamir98/goauth_service/modules/rest"

	"github.com/gin-gonic/gin"
)

// StandardLogin godoc
// @ID standard-login
// @Router /v1/auth/standard/login [POST]
// @Tags login
// @Summary standard login
// @Description standard login
// @Accept json
// @Param platform-id header string true "Platform Id"
// @Param credentials body rest.StandardLoginModel true "credentials"
// @Produce json
// @Success 200 {object} rest.ResponseModel{data=rest.LoginResponseModel} "Success"
// @Response 422 {object} rest.ResponseModel{error=string} "Validation Error"
// @Response 400 {object} rest.ResponseModel{error=string} "Bad Request"
// @Response 401 {object} rest.ResponseModel{error=string} "Unauthorized"
// @Response 403 {object} rest.ResponseModel{error=string} "Forbidden"
// @Failure 500 {object} rest.ResponseModel{error=string} "Server Error"
func (h *Handler) StandardLogin(c *gin.Context) {
	var (
		entity   rest.StandardLoginModel
		response rest.LoginResponseModel
	)

	clientPlatformID := c.GetHeader("platform-id")

	if !util.IsValidUUID(clientPlatformID) {
		h.handleErrorResponse(c, 422, "validation error", "platform-id")
		return
	}

	err := c.ShouldBindJSON(&entity)

	if err != nil {
		h.handleErrorResponse(c, 400, "parse error", err.Error())
		return
	}

	if len(entity.Username) < 6 {
		h.handleErrorResponse(c, 422, "validation error", "username")
		return
	}

	if len(entity.Password) < 6 {
		h.handleErrorResponse(c, 422, "validation error", "password")
		return
	}

	userID, err := h.storageCassandra.Auth().GetUserIDByUsername(entity.Username)
	if err != nil {
		h.handleErrorResponse(c, 401, "unauthorized", "username or password is wrong")
		return
	}

	response.UserFound = true

	user, err := h.storageCassandra.User().GetByID(userID)
	if err != nil {
		h.handleErrorResponse(c, 500, "database error", err.Error())
		return
	}

	response.User = user

	clientType, err := h.storageCassandra.ClientType().GetByID(user.ClientTypeID)
	if err != nil {
		h.handleErrorResponse(c, 500, "database error", err.Error())
		return
	}

	response.ClientType = clientType

	if clientType.LoginStrategy != "STANDARD" {
		h.handleErrorResponse(c, 403, "wrong login strategy", nil)
		return
	}

	if user.Active < 0 {
		h.handleErrorResponse(c, 403, "forbidden", "user is not active")
		return
	}
	if user.Active == 0 {
		h.handleErrorResponse(c, 403, "forbidden", "user hasn't been activated")
		return
	}
	if user.ExpiresAt.Unix() < time.Now().Unix() {
		h.handleErrorResponse(c, 403, "forbidden", "user has been expired")
		return
	}

	if user.Password == "" {
		h.handleErrorResponse(c, 400, "bad request", "user haven't set any password before")
		return
	}

	match, err := security.ComparePassword(user.Password, entity.Password)
	if err != nil {
		h.handleErrorResponse(c, 500, "server error", err.Error())
		return
	}

	if !match {
		h.handleErrorResponse(c, 401, "unauthorized", "username or password is wrong")
		return
	}

	_, err = h.storageCassandra.Client().Get(clientPlatformID, user.ClientTypeID)
	if err != nil {
		h.handleErrorResponse(c, 403, "forbidden", "wrong platform id")
		return
	}

	sessions, err := h.storageCassandra.Session().GetByUserID(user.ID)
	if err != nil {
		h.handleErrorResponse(c, 500, "database error", err.Error())
		return
	}

	response.UserSessions = sessions

	uuid, err := uuid.NewRandom()
	if err != nil {
		h.handleErrorResponse(c, 500, "server error", err.Error())
		return
	}

	session := rest.SessionModel{
		ClientPlatformID: clientPlatformID,
		ClientTypeID:     user.ClientTypeID,
		UserID:           user.ID,
		ID:               uuid.String(),
		RoleID:           user.RoleID,
		IP:               c.ClientIP(),
		Data:             "",
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
		ExpiresAt:        time.Now().Add(config.RtExpireInTime),
	}

	err = h.storageCassandra.Session().Create(session)
	if err != nil {
		h.handleErrorResponse(c, 500, "database error", err.Error())
		return
	}

	m := map[string]interface{}{
		"client_platform_id": session.ClientPlatformID,
		"client_type_id":     session.ClientTypeID,
		"user_id":            session.UserID,
		"id":                 session.ID,
	}

	accessToken, err := security.GenerateJWT(m, config.AtExpireInTime, h.cfg.SecretKey)
	if err != nil {
		h.handleErrorResponse(c, 500, "server error", err.Error())
		return
	}

	refreshToken, err := security.GenerateJWT(m, config.AtExpireInTime, h.cfg.SecretKey)
	if err != nil {
		h.handleErrorResponse(c, 500, "server error", err.Error())
		return
	}

	response.Token = rest.TokenModel{
		AccessToken:      accessToken,
		RefreshToken:     refreshToken,
		CreatedAt:        session.CreatedAt,
		UpdatedAt:        session.ExpiresAt,
		ExpiresAt:        session.ExpiresAt,
		RefreshInSeconds: int(config.AtExpireInTime.Seconds()),
	}

	h.handleSuccessResponse(c, 200, "ok", response)
}
