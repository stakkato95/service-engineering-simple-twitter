package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stakkato95/twitter-service-analytics/dto"
	"github.com/stakkato95/twitter-service-analytics/service"
)

type AnalyticsHandler struct {
	service service.TweetService
}

func (h *AnalyticsHandler) getAnalytics(ctx *gin.Context) {
	count := h.service.GetTweetCount()
	ctx.JSON(http.StatusOK, dto.ResponseDto{Data: dto.TweetCountDto{TweetCount: count}})
}

func errorResponse(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, dto.ResponseDto{Error: err.Error()})
}
