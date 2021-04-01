package v1

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/saidamir98/goauth_service/modules/rest"
	"github.com/saidamir98/goauth_service/pkg/security"
	"github.com/saidamir98/goauth_service/pkg/util"
)

// CreateUser godoc
// @ID create-user
// @Router /v1/auth/user [POST]
// @Tags user
// @Summary create user
// @Description create user
// @Accept json
// @Param input body rest.CreateUserModel true "body"
// @Produce json
// @Success 200 {object} rest.ResponseModel{data=string} "Success"
// @Response 422 {object} rest.ResponseModel{error=string} "Validation Error"
// @Response 400 {object} rest.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} rest.ResponseModel{error=string} "Server Error"
func (h *Handler) CreateUser(c *gin.Context) {
	var (
		entity rest.CreateUserModel
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

	fmt.Println(clientType)
	//
	// TODO - validate entity by clientType rule
	//

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

	userID, err := h.storageCassandra.User().Create(entity)
	if err != nil {
		h.handleErrorResponse(c, 400, "database error", err.Error())
		return
	}

	entity.ID = userID

	h.handleSuccessResponse(c, 200, "ok", entity.ID)
}

// UpdateUser godoc
// @ID update-user
// @Router /v1/auth/user [PUT]
// @Tags user
// @Summary update user
// @Description update user
// @Accept json
// @Param input body rest.UpdateUserModel true "body"
// @Produce json
// @Success 200 {object} rest.ResponseModel{data=string} "Success"
// @Response 422 {object} rest.ResponseModel{error=string} "Validation Error"
// @Response 400 {object} rest.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} rest.ResponseModel{error=string} "Server Error"
func (h *Handler) UpdateUser(c *gin.Context) {
	var (
		entity rest.UpdateUserModel
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

	if !util.IsValidUUID(entity.RoleID) {
		h.handleErrorResponse(c, 422, "validation error", "role_id")
		return
	}

	err = h.storageCassandra.User().Update(entity)
	if err != nil {
		h.handleErrorResponse(c, 400, "database error", err.Error())
		return
	}

	h.handleSuccessResponse(c, 200, "ok", entity.ID)
}
