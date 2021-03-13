package api

import (
	v1 "github.com/saidamir98/goauth_service/api/handlers/v1"
	"github.com/saidamir98/goauth_service/config"
	"github.com/saidamir98/goauth_service/pkg/cors"
	"github.com/saidamir98/goauth_service/pkg/logger"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// New ...
// @title Go Boilerplate API
// @version 1.0
// @description This is a Go Boilerplate for medium sized projects
// @contact.name Saidamir Botirov
// @contact.email saidamir.botirov@gmail.com
// @contact.url https://www.linkedin.com/in/saidamir-botirov-a08559192
func New(cfg config.Config, log logger.Logger) (*gin.Engine, error) {
	if cfg.Environment != "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	router.Use(gin.Logger(), gin.Recovery()) // Later they will be replaced by custom Logger and Recovery

	router.Use(cors.MyCORSMiddleware())

	handlerV1 := v1.New(cfg, log)

	router.GET("/ping", handlerV1.Ping)
	router.GET("/config", handlerV1.GetConfig)

	rV1 := router.Group("/v1")
	{
		endpointsV1(rV1, handlerV1)
	}

	url := ginSwagger.URL("/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router, nil
}
