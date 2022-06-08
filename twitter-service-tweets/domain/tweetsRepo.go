package domain

import (
	"fmt"

	"github.com/stakkato95/service-engineering-go-lib/logger"
)

type TweetsRepo interface {
	AddTweet(Tweet) *Tweet
	GetAllTweets(int) []Tweet
}

type simpleTweetsRepo struct {
	repo DbRepo
}

func NewTweetsRepo(repo DbRepo) TweetsRepo {
	return &simpleTweetsRepo{repo}
}

func (r *simpleTweetsRepo) AddTweet(tweet Tweet) *Tweet {
	r.repo.GetDb().Create(&tweet)
	logger.Info(fmt.Sprintf("added tweet with id: %d", tweet.Id))
	return &tweet
}

func (r *simpleTweetsRepo) GetAllTweets(userId int) []Tweet {
	tweets := []Tweet{}
	r.repo.GetDb().Where("user_id = ?", userId).Find(&tweets)
	return tweets
}
