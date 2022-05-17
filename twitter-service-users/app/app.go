package app

import (
	"github.com/stakkato95/service-engineering-go-lib/logger"
	"github.com/stakkato95/twitter-service-users/config"
)

func Start() {
	logger.Info("works")
	logger.Info(config.AppConfig.ServerPort)
}
