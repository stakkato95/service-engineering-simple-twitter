package app

import (
	"github.com/gin-gonic/gin"
	"github.com/stakkato95/twitter-service-tweets/config"
	"github.com/stakkato95/twitter-service-tweets/domain"
	"github.com/stakkato95/twitter-service-tweets/service"
)

func Start() {
	repo := domain.NewTweetsRepo()
	sink := domain.NewTweetsSink()
	service := service.NewTweetsService(repo, sink)

	h := TweetsHandler{service}

	router := gin.Default()
	router.POST("/tweets", h.addTweet)
	router.GET("/tweets/:userId", h.getTweets)
	router.Run(config.AppConfig.ServerPort)
}
