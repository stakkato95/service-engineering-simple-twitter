package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/stakkato95/service-engineering-go-lib/logger"
	"github.com/stakkato95/twitter-service-users/config"
	"github.com/stakkato95/twitter-service-users/service"
)

func Start() {
	logger.Info("works")
	logger.Info(config.AppConfig.ServerPort)

	service := service.NewUserService()
	h := userHandlers{service}

	router := chi.NewRouter()

	router.Get("/hello", h.hello)

	logger.Info("users service listening on port " + config.AppConfig.ServerPort)
	logger.Fatal("ca not run server " + http.ListenAndServe(config.AppConfig.ServerPort, router).Error())
}
