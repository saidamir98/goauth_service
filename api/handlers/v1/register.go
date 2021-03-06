package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/saidamir98/goauth_service/modules/rest"
	"github.com/saidamir98/goauth_service/pkg/security"
	"github.com/saidamir98/goauth_service/pkg/util"
)

// RegisterUser godoc
// @ID register-user
// @Router /v1/auth/user/register [POST]
// @Tags register
// @Summary register user
// @Description register user
// @Accept json
// @Param input body rest.RegisterUserModel true "body"
// @Produce json
// @Success 200 {object} rest.ResponseModel{data=string} "Success"
// @Response 422 {object} rest.ResponseModel{error=string} "Validation Error"
// @Response 400 {object} rest.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} rest.ResponseModel{error=string} "Server Error"
func (h *Handler) RegisterUser(c *gin.Context) {
	var (
		entity rest.RegisterUserModel
	)
	err := c.ShouldBindJSON(&entity)

	if err != nil {
		h.handleErrorResponse(c, 400, "parse error", err)
		return
	}

	if !util.IsValidUUID(entity.ClientTypeID) {
		h.handleErrorResponse(c, 422, "validation error", "client_type_id")
		return
	}

	clientType, err := h.storageCassandra.ClientType().GetByID(entity.ClientTypeID)
	if err != nil {
		h.handleErrorResponse(c, 500, "database error", err.Error())
		return
	}

	//
	// TODO - validate entity by clientType rule
	//

	if !clientType.SelfRegister {
		h.handleErrorResponse(c, 400, "bad request", "this client type cannot self register")
		return
	}

	if !util.IsValidUUID(entity.RoleID) {
		h.handleErrorResponse(c, 422, "validation error", "role_id")
		return
	}

	if !util.IsValidEmail(entity.Email) {
		h.handleErrorResponse(c, 422, "validation error", "email")
		return
	}

	if !util.IsValidLogin(entity.Login) {
		h.handleErrorResponse(c, 422, "validation error", "login")
		return
	}

	for i := 0; i < len(entity.Phones); i++ {
		if !util.IsValidPhone(entity.Phones[i]) {
			h.handleErrorResponse(c, 422, "validation error", "phone")
			return
		}
	}

	if len(entity.Password) < 6 {
		h.handleErrorResponse(c, 422, "validation error", "password")
		return
	}

	hashedPassword, err := security.HashPassword(entity.Password)
	if err != nil {
		h.handleErrorResponse(c, 500, "server error", err.Error())
		return
	}

	entity.Password = hashedPassword
	entity.Active = 0

	userID, err := h.storageCassandra.Auth().RegisterUser(entity)
	if err != nil {
		h.handleErrorResponse(c, 400, "database error", err.Error())
		return
	}

	entity.ID = userID

	// clientType.ConfirmBy
	//
	// TODO  - generate and send activation passcode and token
	//

	h.handleSuccessResponse(c, 200, "ok", entity.ID)
}
