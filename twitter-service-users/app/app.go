package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/stakkato95/service-engineering-go-lib/logger"
	"github.com/stakkato95/twitter-service-users/config"
	"github.com/stakkato95/twitter-service-users/domain"
	"github.com/stakkato95/twitter-service-users/service"
)

func Start() {
	repo := domain.NewUserRepo()
	service := service.NewUserService(repo)
	h := userHandlers{service}

	router := chi.NewRouter()

	//chi nested handlers
	router.Get("/debug/hello", h.hello)
	router.Post("/debug/create", h.create)
	router.Post("/debug/auth", h.auth)

	logger.Info("users service listening on port " + config.AppConfig.ServerPort)
	logger.Fatal("can not run server: " + http.ListenAndServe(config.AppConfig.ServerPort, router).Error())
}
