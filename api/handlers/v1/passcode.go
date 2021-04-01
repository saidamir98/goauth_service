package v1

import (
	"fmt"
	"time"

	"github.com/saidamir98/goauth_service/config"
	"github.com/saidamir98/goauth_service/pkg/security"
	"github.com/saidamir98/goauth_service/pkg/util"

	"github.com/google/uuid"
	"github.com/saidamir98/goauth_service/modules/rest"

	"github.com/gin-gonic/gin"
)

// GeneratePasscode godoc
// @ID generate-passcode
// @Router /v1/auth/passcode/generate [POST]
// @Tags passcode
// @Summary generate passcode
// @Description generate passcode
// @Accept json
// @Param platform-id header string true "Platform Id"
// @Param credentials body rest.GeneratePasscodeModel true "credentials"
// @Produce json
// @Success 200 {object} rest.ResponseModel{data=rest.GeneratePasscodeResponseModel} "Success"
// @Response 422 {object} rest.ResponseModel{error=string} "Validation Error"
// @Response 400 {object} rest.ResponseModel{error=string} "Bad Request"
// @Response 401 {object} rest.ResponseModel{error=string} "Unauthorized"
// @Response 403 {object} rest.ResponseModel{error=string} "Forbidden"
// @Failure 500 {object} rest.ResponseModel{error=string} "Server Error"
func (h *Handler) GeneratePasscode(c *gin.Context) {
	var (
		entity   rest.GeneratePasscodeModel
		response rest.GeneratePasscodeResponseModel
	)

	response.Period = 60

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

	clientType, err := h.storageCassandra.ClientType().GetByID(entity.ClientTypeID)
	if err != nil {
		h.handleErrorResponse(c, 500, "database error", err.Error())
		return
	}

	response.ClientType = clientType

	switch clientType.ConfirmBy {
	case "EMAIL":
		if !util.IsValidEmail(entity.Username) {
			h.handleErrorResponse(c, 422, "validation error", "username")
			return
		}
	case "PHONE":
		if !util.IsValidPhone(entity.Username) {
			h.handleErrorResponse(c, 422, "validation error", "username")
			return
		}
	default:
		h.handleErrorResponse(c, 422, "validation error", "username")
		return
	}

	userID, err := h.storageCassandra.Auth().GetUserIDByUsername(entity.Username)
	if err != nil {
		response.UserFound = false
		response.Period = response.Period * 10
		//
		// TODO - if self_register is enabled generate and send passcode to register
		//
		h.handleErrorResponse(c, 200, "ok", response)
		return
	}
	response.UserFound = true

	user, err := h.storageCassandra.User().GetByID(userID)
	if err != nil {
		h.handleErrorResponse(c, 500, "database error", err.Error())
		return
	}

	response.User = user

	if user.ClientTypeID != clientType.ID {
		h.handleErrorResponse(c, 403, "forbidden", "mismatch between given client_type_id and user client_type_id")
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

	switch clientType.LoginStrategy { // diffrencec between "PASSCODE" and "OTP" strategies that the first one compares password to be successful
	case "PASSCODE":
		if len(entity.Password) < 6 {
			h.handleErrorResponse(c, 422, "validation error", "password")
			return
		}

		if user.Password == "" {
			h.handleErrorResponse(c, 401, "bad request", "user haven't set any password before")
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
	case "OTP":

	default:
		h.handleErrorResponse(c, 400, "bad request", "wrong login strategy")
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

	code, err := security.GenerateRandomStringByPool(h.cfg.PasscodeLength, h.cfg.PasscodePool)
	if err != nil {
		h.handleErrorResponse(c, 500, "server error", err.Error())
		return
	}

	hashedCode, err := security.HashPassword(code)
	if err != nil {
		h.handleErrorResponse(c, 500, "server error", err.Error())
		return
	}

	uuid, err := uuid.NewRandom()
	if err != nil {
		h.handleErrorResponse(c, 500, "server error", err.Error())
		return
	}

	duration := time.Duration(response.Period) * time.Second

	passcode := rest.PasscodeModel{
		UserID:     user.ID,
		ID:         uuid.String(),
		ConfirmBy:  clientType.ConfirmBy,
		HashedCode: hashedCode,
		State:      0,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		ExpiresAt:  time.Now().Add(duration),
	}

	err = h.storageCassandra.Passcode().Create(passcode)
	if err != nil {
		h.handleErrorResponse(c, 500, "database error", err.Error())
		return
	}

	m := map[string]interface{}{
		"user_id": passcode.UserID,
		"id":      passcode.ID,
	}

	passcodeToken, err := security.GenerateJWT(m, duration, h.cfg.SecretKey)
	if err != nil {
		h.handleErrorResponse(c, 500, "server error", err.Error())
		return
	}
	response.PasscodeToken = passcodeToken

	fmt.Println(code)
	//
	// TODO - send code to the user email or phone
	//

	h.handleSuccessResponse(c, 200, "ok", response)
	return
}

// ConfirmPasscode godoc
// @ID confirm-passcode
// @Router /v1/auth/passcode/confirm [POST]
// @Tags passcode
// @Summary confirm passcode
// @Description confirm passcode
// @Accept json
// @Param platform-id header string true "Platform Id"
// @Param credentials body rest.ConfirmPasscodeModel true "credentials"
// @Produce json
// @Success 200 {object} rest.ResponseModel{data=rest.LoginResponseModel} "Success"
// @Response 422 {object} rest.ResponseModel{error=string} "Validation Error"
// @Response 400 {object} rest.ResponseModel{error=string} "Bad Request"
// @Response 401 {object} rest.ResponseModel{error=string} "Unauthorized"
// @Response 403 {object} rest.ResponseModel{error=string} "Forbidden"
// @Failure 500 {object} rest.ResponseModel{error=string} "Server Error"
func (h *Handler) ConfirmPasscode(c *gin.Context) {
	var (
		entity   rest.ConfirmPasscodeModel
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

	claims, err := security.ExtractClaims(entity.PasscodeToken, h.cfg.SecretKey)
	if err != nil {
		h.handleErrorResponse(c, 401, "passcode token error", err.Error())
		return
	}

	userID := claims["user_id"].(string)
	id := claims["id"].(string)

	user, err := h.storageCassandra.User().GetByID(userID)
	if err != nil {
		h.handleErrorResponse(c, 500, "database error", err.Error())
		return
	}

	response.UserFound = true
	response.User = user

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

	clientType, err := h.storageCassandra.ClientType().GetByID(user.ClientTypeID)
	if err != nil {
		h.handleErrorResponse(c, 500, "database error", err.Error())
		return
	}

	response.ClientType = clientType

	_, err = h.storageCassandra.Client().Get(clientPlatformID, user.ClientTypeID)
	if err != nil {
		h.handleErrorResponse(c, 403, "forbidden", "wrong platform id")
		return
	}

	passcode, err := h.storageCassandra.Passcode().Get(userID, id)
	if err != nil {
		h.handleErrorResponse(c, 403, "forbidden", err.Error())
		return
	}

	if !(-3 < passcode.State || passcode.State < 1) {
		h.handleErrorResponse(c, 403, "forbidden", fmt.Sprintf("passcode state: %d", passcode.State))
		return
	}

	if !(entity.Passcode == "12344321" && (h.cfg.Environment == "development" || h.cfg.Environment == "staging")) {
		match, err := security.ComparePassword(passcode.HashedCode, entity.Passcode)
		if err != nil {
			h.handleErrorResponse(c, 500, "server error", err.Error())
			return
		}

		passcode.UpdatedAt = time.Now()
		if !match {
			passcode.State = passcode.State - 1
			h.handleErrorResponse(c, 403, "forbidden", "wrong passcode or passcode token")
			//
			// TODO - Update passcode
			//
			return
		}

		passcode.State = 1
		//
		// TODO - Update passcode
		//
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
	return

}
