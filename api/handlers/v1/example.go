package v1

import (
	"github.com/saidamir98/goauth/modules/rest"
	"github.com/saidamir98/goauth_service/pkg/logger"

	"github.com/gin-gonic/gin"
)

// Ping godoc
// @ID ping
// @Router /ping [GET]
// @Summary returns "pong" message
// @Description this returns "pong" messsage to show service is working
// @Accept json
// @Produce json
// @Success 200 {object} rest.ResponseModel{data=string} "desc"
// @Failure 500 {object} rest.ResponseModel{error=string}
func (h *Handler) Ping(c *gin.Context) {
	c.JSON(200, rest.ResponseModel{
		Code:    200,
		Message: "ok",
		Data:    "pong",
	})
	return
}

// GetConfig godoc
// @ID get-config
// @Router /config [GET]
// @Summary gets project config
// @Description shows config of the project only on the development phase
// @Accept json
// @Produce json
// @Success 200 {object} rest.ResponseModel{data=config.Config} "desc"
// @Response 400 {object} rest.ResponseModel{error=string} "Bad Request"
func (h *Handler) GetConfig(c *gin.Context) {
	h.log.Info("get config", logger.Any("result", h.cfg))
	switch h.cfg.Environment {
	case "development":
		h.handleSuccessResponse(c, 200, "ok", h.cfg)
		return
	case "staging":
		h.handleSuccessResponse(c, 200, h.cfg.Environment, nil)
		return
	case "production":
		h.handleSuccessResponse(c, 200, "private data", nil)
		return
	}

	h.handleErrorResponse(c, 400, "wrong environment", h.cfg.Environment)
	return
}
