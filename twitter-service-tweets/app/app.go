package app

import (
	"github.com/gin-gonic/gin"
	"github.com/stakkato95/twitter-service-tweets/config"
	"github.com/stakkato95/twitter-service-tweets/domain"
	"github.com/stakkato95/twitter-service-tweets/service"
)

func Start() {
	db := domain.NewDbRepo()

	subscriptionRepo := domain.NewSubscriptionRepo(db)
	subscriptionService := service.NewSubscriptionService(subscriptionRepo)

	tweetsRepo := domain.NewTweetsRepo(db)
	tweetsSink := domain.NewTweetsSink()
	tweetsService := service.NewTweetsService(tweetsRepo, tweetsSink)

	h := TweetsHandler{
		tweetsService:       tweetsService,
		subscriptionService: subscriptionService,
	}

	router := gin.Default()

	router.POST("/tweets", h.addTweet)
	router.GET("/tweets/:userId", h.getTweets)

	router.POST("/subscription", h.addSubscription)

	router.Run(config.AppConfig.ServerPort)
}
