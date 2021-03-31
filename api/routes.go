package api

import (
	v1 "github.com/saidamir98/goauth_service/api/handlers/v1"

	"github.com/gin-gonic/gin"
)

func endpointsV1(r *gin.RouterGroup, h *v1.Handler) {
	r.POST("/auth/standard/login", h.StandardLogin)

	r.POST("/auth/has-access", h.HasAccess)

	r.POST("/auth/passcode/generate", h.GeneratePasscode)
	r.POST("/auth/passcode/confirm", h.ConfirmPasscode)

	r.POST("/auth/user/register", h.RegisterUser)
	// r.POST("/auth/user/activate", h.ActivateUser)
	// r.POST("/auth/user/deactivate", h.DeactivateUser)

	r.POST("/auth/user", h.CreateUser)
	r.PUT("/auth/user", h.UpdateUser)
	// r.PUT("/auth/user/password", h.UpdateUserPassword)
	// r.PUT("/auth/user/login", h.UpdateUserLogin)
	// r.PUT("/auth/user/email", h.UpdateUserEmail)
	// r.PUT("/auth/user/phone", h.UpdateUserPhone)
	// r.GET("/auth/user", h.GetUserList)
	// r.DELETE("/auth/user", h.DeleteUser)

	r.PUT("/auth/refresh", h.RefreshToken)
	r.DELETE("/auth/logout", h.Logout)
}
