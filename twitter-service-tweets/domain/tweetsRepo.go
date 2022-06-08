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
	user := User{Id: userId}
	r.repo.GetDb().Preload("Subscriptions").First(&user)

	tweets := []Tweet{}
	for _, subscription := range user.Subscriptions {
		tmp := []Tweet{}
		r.repo.GetDb().Model(&Tweet{}).Where("user_id = ?", subscription.To).Find(&tmp)
		tweets = append(tweets, tmp...)
	}
	return tweets
}
