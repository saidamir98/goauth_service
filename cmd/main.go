package main

import (
	"github.com/saidamir98/goauth_service/api"
	"github.com/saidamir98/goauth_service/api/docs"
	"github.com/saidamir98/goauth_service/config"
	"github.com/saidamir98/goauth_service/pkg/logger"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.App, cfg.LogLevel)

	docs.SwaggerInfo.Host = cfg.ServiceHost + cfg.HTTPPort
	// docs.SwaggerInfo.BasePath = cfg.BasePath
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	apiServer, err := api.New(cfg, log)
	if err != nil {
		log.Panic("error on the api server", logger.Error(err))
	}

	err = apiServer.Run(cfg.HTTPPort)
	if err != nil {
		log.Panic("error on the api server", logger.Error(err))
	}

	log.Panic("api server has finished")
}
