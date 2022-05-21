package app

import (
	"github.com/gin-gonic/gin"
	"github.com/stakkato95/twitter-service-tweets/config"
	"github.com/stakkato95/twitter-service-tweets/domain"
	"github.com/stakkato95/twitter-service-tweets/service"
)

func Start() {
	repo := domain.NewTweetsRepo()
	service := service.NewTweetsService(repo)

	h := TweetsHandler{service}

	router := gin.Default()
	router.POST("/tweets", h.addTweet)
	router.Run(config.AppConfig.ServerPort)
}
