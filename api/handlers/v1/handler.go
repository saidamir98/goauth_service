package v1

import (
	"strconv"

	"github.com/saidamir98/goauth_service/config"
	"github.com/saidamir98/goauth_service/modules/rest"
	"github.com/saidamir98/goauth_service/pkg/logger"
	"github.com/saidamir98/goauth_service/storage"

	"github.com/gin-gonic/gin"
)

// Handler ...
type Handler struct {
	cfg              config.Config
	log              logger.Logger
	storageCassandra storage.CassandraStorageI
}

// New ...
func New(cfg config.Config, log logger.Logger) *Handler {
	return &Handler{
		cfg:              cfg,
		log:              log,
		storageCassandra: storage.NewStorageCassandra(cfg),
	}
}

func (h *Handler) handleSuccessResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, rest.ResponseModel{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func (h *Handler) handleErrorResponse(c *gin.Context, code int, message string, err interface{}) {
	h.log.Error(message, logger.Int("code", code), logger.Any("error", err))
	c.JSON(code, rest.ResponseModel{
		Code:    code,
		Message: message,
		Error:   err,
	})
}

func (h *Handler) parseOffsetQueryParam(c *gin.Context) (int, error) {
	return strconv.Atoi(c.DefaultQuery("offset", h.cfg.DefaultOffset))
}

func (h *Handler) parseLimitQueryParam(c *gin.Context) (int, error) {
	return strconv.Atoi(c.DefaultQuery("limit", h.cfg.DefaultLimit))
}
