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
	// r.PUT("/auth/user/:id", h.UpdateUser)
	// r.PUT("/auth/user/password", h.UpdateUserPassword)
	// r.DELETE("/auth/user", h.DeleteUser)

	r.PUT("/auth/refresh", h.RefreshToken)
	r.DELETE("/auth/logout", h.Logout)
}
