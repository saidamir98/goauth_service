package v1

import (
	"time"

	"github.com/saidamir98/goauth_service/config"
	"github.com/saidamir98/goauth_service/pkg/logger"
	"github.com/saidamir98/goauth_service/pkg/security"
	"github.com/saidamir98/goauth_service/pkg/util"

	"github.com/saidamir98/goauth_service/modules/rest"

	"github.com/gin-gonic/gin"
)

// HasAccess godoc
// @ID has-access
// @Router /v1/auth/has-access [POST]
// @Tags auth
// @Summary has access
// @Description has access
// @Accept json
// @Param platform-id header string true "Platform Id"
// @Param access body rest.AccessModel true "Access Info"
// @Produce json
// @Success 200 {object} rest.ResponseModel{data=rest.HasAccessModel} "Success"
// @Response 422 {object} rest.ResponseModel{error=string} "Validation Error"
// @Response 400 {object} rest.ResponseModel "Bad Request"
// @Response 401 {object} rest.ResponseModel{error=string} "Unauthorized"
// @Response 403 {object} rest.ResponseModel{error=string} "Forbidden"
// @Failure 500 {object} rest.ResponseModel{error=string} "Server Error"
func (h *Handler) HasAccess(c *gin.Context) {
	var (
		entity rest.AccessModel
	)

	_clientPlatformID := c.GetHeader("platform-id")

	if !util.IsValidUUID(_clientPlatformID) {
		h.handleErrorResponse(c, 422, "validation error", "platform-id")
		return
	}

	err := c.ShouldBindJSON(&entity)

	if err != nil {
		h.handleErrorResponse(c, 400, "parse error", err.Error())
		return
	}

	claims, err := security.ExtractClaims(entity.AccessToken, h.cfg.SecretKey)
	if err != nil {
		h.handleErrorResponse(c, 401, "token error", err.Error())
		return
	}

	clientPlatformID := claims["client_platform_id"].(string)
	clientTypeID := claims["client_type_id"].(string)
	userID := claims["user_id"].(string)
	id := claims["id"].(string)

	if _clientPlatformID != clientPlatformID {
		h.handleErrorResponse(c, 401, "unauthorized", "mismatch platform-id with token info")
	}

	session, err := h.storageCassandra.Session().Get(clientPlatformID, clientTypeID, userID, id)
	if err != nil {
		h.handleErrorResponse(c, 403, "forbidden", "session not found")
		return
	}

	_, err = h.storageCassandra.Client().Get(clientPlatformID, session.ClientTypeID)
	if err != nil {
		h.handleErrorResponse(c, 403, "forbidden", "platform blocked")
		return
	}

	user, err := h.storageCassandra.User().GetByID(session.UserID)
	if err != nil {
		h.handleErrorResponse(c, 500, "database error", err.Error())
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

	_, err = h.storageCassandra.CachedData().RoleScopePermission(user.RoleID, entity.Method, entity.URL)
	if err != nil {
		h.handleErrorResponse(c, 401, "unauthorized", "access denied")
		return
	}

	h.handleSuccessResponse(c, 200, "ok", rest.HasAccessModel{
		ClientPlatformID: session.ClientPlatformID,
		ClientTypeID:     session.ClientTypeID,
		UserID:           session.UserID,
		ID:               session.ID,
		RoleID:           session.RoleID,
	})
}

// RefreshToken godoc
// @ID refresh-token
// @Router /v1/auth/refresh [PUT]
// @Tags auth
// @Summary refresh user token
// @Description refresh user token
// @Accept json
// @Param platform-id header string true "Platform Id"
// @Param token body rest.RefreshTokenModel true "Token Info"
// @Produce json
// @Success 200 {object} rest.ResponseModel{data=rest.TokenModel} "Success"
// @Response 422 {object} rest.ResponseModel{error=string} "Validation Error"
// @Response 400 {object} rest.ResponseModel{error=string} "Bad Request"
// @Response 401 {object} rest.ResponseModel{error=string} "Unauthorized"
// @Response 403 {object} rest.ResponseModel{error=string} "Forbidden"
// @Failure 500 {object} rest.ResponseModel{error=string} "Server Error"
func (h *Handler) RefreshToken(c *gin.Context) {
	var (
		entity rest.RefreshTokenModel
	)

	_clientPlatformID := c.GetHeader("platform-id")

	if !util.IsValidUUID(_clientPlatformID) {
		h.handleErrorResponse(c, 422, "validation error", "platform-id")
		return
	}

	err := c.ShouldBindJSON(&entity)

	if err != nil {
		h.handleErrorResponse(c, 400, "parse error", err.Error())
		return
	}

	claims, err := security.ExtractClaims(entity.RefreshToken, h.cfg.SecretKey)
	if err != nil {
		h.handleErrorResponse(c, 401, "token error", err.Error())
		return
	}

	clientPlatformID := claims["client_platform_id"].(string)
	clientTypeID := claims["client_type_id"].(string)
	userID := claims["user_id"].(string)
	id := claims["id"].(string)

	if _clientPlatformID != clientPlatformID {
		h.handleErrorResponse(c, 401, "unauthorized", "mismatch platform-id with token info")
		return
	}

	session, err := h.storageCassandra.Session().Get(clientPlatformID, clientTypeID, userID, id)
	if err != nil {
		h.handleErrorResponse(c, 403, "forbidden", err.Error())
		return
	}

	_, err = h.storageCassandra.Client().Get(clientPlatformID, session.ClientTypeID)
	if err != nil {
		h.handleErrorResponse(c, 403, "forbidden", "platform blocked")
		return
	}

	user, err := h.storageCassandra.User().GetByID(session.UserID)
	if err != nil {
		h.handleErrorResponse(c, 500, "database error", err.Error())
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

	if session.ClientTypeID != user.ClientTypeID {
		h.handleErrorResponse(c, 403, "forbidden", "user type has been updated")
		return
	}

	session.RoleID = user.RoleID
	session.IP = c.ClientIP()
	session.UpdatedAt = time.Now()
	session.ExpiresAt = time.Now().Add(config.RtExpireInTime)

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

	h.handleSuccessResponse(c, 200, "ok", rest.TokenModel{
		AccessToken:      accessToken,
		RefreshToken:     refreshToken,
		CreatedAt:        session.CreatedAt,
		UpdatedAt:        session.ExpiresAt,
		ExpiresAt:        session.ExpiresAt,
		RefreshInSeconds: int(config.AtExpireInTime.Seconds()),
	})
}

// Logout godoc
// @ID logout
// @Router /v1/auth/logout [DELETE]
// @Tags auth
// @Summary logout user
// @Description logout user by his/her token
// @Accept json
// @Param platform-id header string true "Platform Id"
// @Param Authorization header string true "Bearer Token"
// @Produce json
// @Success 204 {object} rest.ResponseModel{data=string} "Success"
// @Response 422 {object} rest.ResponseModel{error=string} "Validation Error"
// @Response 400 {object} rest.ResponseModel{error=string} "Bad Request"
// @Response 401 {object} rest.ResponseModel{error=string} "Unauthorized"
// @Response 403 {object} rest.ResponseModel{error=string} "Forbidden"
// @Failure 500 {object} rest.ResponseModel{error=string} "Server Error"
func (h *Handler) Logout(c *gin.Context) {
	_clientPlatformID := c.GetHeader("platform-id")

	if !util.IsValidUUID(_clientPlatformID) {
		h.handleErrorResponse(c, 422, "validation error", "platform-id")
		return
	}

	token, err := security.ExtractToken(c.GetHeader("Authorization"))

	if err != nil {
		h.handleErrorResponse(c, 400, "validation error", err.Error())
		return
	}

	claims, err := security.ExtractClaims(token, h.cfg.SecretKey)
	if err != nil {
		h.handleErrorResponse(c, 401, "token error", err.Error())
		return
	}

	clientPlatformID := claims["client_platform_id"].(string)
	clientTypeID := claims["client_type_id"].(string)
	userID := claims["user_id"].(string)
	id := claims["id"].(string)

	if _clientPlatformID != clientPlatformID {
		h.handleErrorResponse(c, 401, "unauthorized", "mismatch platform-id with token info")
		return
	}

	err = h.storageCassandra.Session().Delete(clientPlatformID, clientTypeID, userID, id)
	if err != nil {
		h.handleErrorResponse(c, 403, "forbidden", err.Error())
		return
	}

	h.log.Info("logout",
		logger.String("client_platform_id", clientPlatformID),
		logger.String("client_type_id", clientTypeID),
		logger.String("user_id", userID),
		logger.String("id", id),
	)

	h.handleSuccessResponse(c, 204, "", nil) // status 204 returns no content
}
