package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stakkato95/twitter-service-tweets/dto"
	"github.com/stakkato95/twitter-service-tweets/service"
)

type getTweetsUriParams struct {
	UserId int `uri:"userId" binding:"required,min=1"`
}

type TweetsHandler struct {
	tweetsService       service.TweetsService
	subscriptionService service.SubscriptionService
}

func (h *TweetsHandler) addTweet(ctx *gin.Context) {
	var tweetDto dto.TweetDto
	if err := ctx.ShouldBindJSON(&tweetDto); err != nil {
		errorResponse(ctx, err)
		return
	}

	createdTweet := h.tweetsService.AddTweet(tweetDto)
	ctx.JSON(http.StatusOK, dto.ResponseDto{Data: *createdTweet})
}

func (h *TweetsHandler) getTweets(ctx *gin.Context) {
	var uriParams getTweetsUriParams
	if err := ctx.ShouldBindUri(&uriParams); err != nil {
		errorResponse(ctx, err)
		return
	}

	tweets := h.tweetsService.GetAllTweets(uriParams.UserId)
	ctx.JSON(http.StatusOK, dto.ResponseDto{Data: tweets})
}

func (h *TweetsHandler) addSubscription(ctx *gin.Context) {
	var subscriptionDto dto.SubscriptionDto
	if err := ctx.ShouldBindJSON(&subscriptionDto); err != nil {
		errorResponse(ctx, err)
		return
	}

	h.subscriptionService.AddSubscription(subscriptionDto)
	ctx.JSON(http.StatusOK, dto.ResponseDto{Data: "ok"})
}

func errorResponse(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, dto.ResponseDto{Error: err.Error()})
}
