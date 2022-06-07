package app

import (
	"github.com/gin-gonic/gin"
	"github.com/stakkato95/twitter-service-analytics/config"
	"github.com/stakkato95/twitter-service-analytics/domain"
	"github.com/stakkato95/twitter-service-analytics/service"
)

func Start() {
	repo := domain.NewTweetProcessor()
	service := service.NewTweetService(repo)

	h := AnalyticsHandler{service: service}

	router := gin.Default()
	router.GET("/analytics", h.getAnalytics)
	router.Run(config.AppConfig.ServerPort)
}
