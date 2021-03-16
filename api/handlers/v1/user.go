package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/saidamir98/goauth/modules/rest"
	"github.com/saidamir98/goauth_service/pkg/security"
	"github.com/saidamir98/goauth_service/pkg/util"
)

// RegisterUser godoc
// @ID register-user
// @Router /v1/auth/user/register [POST]
// @Tags auth
// @Summary register user
// @Description register user
// @Accept json
// @Param input body rest.RegisterUserModel true "body"
// @Produce json
// @Success 200 {object} rest.ResponseModel{data=string} "Success"
// @Response 422 {object} rest.ResponseModel{error=string} "Validation Error"
// @Response 400 {object} rest.ResponseModel "Bad Request"
// @Failure 500 {object} rest.ResponseModel "Server Error"
func (h *Handler) RegisterUser(c *gin.Context) {
	var (
		entity rest.RegisterUserModel
	)
	err := c.ShouldBindJSON(&entity)

	if err != nil {
		h.handleErrorResponse(c, 400, "parse error", err)
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

	userID, err := h.storageCassandra.User().RegisterUser(entity)
	if err != nil {
		h.handleErrorResponse(c, 400, "database error", err.Error())
		return
	}

	entity.ID = userID

	h.handleSuccessResponse(c, 200, "ok", entity.ID)
}
