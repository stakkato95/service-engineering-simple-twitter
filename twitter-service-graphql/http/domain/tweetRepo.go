package domain

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/stakkato95/service-engineering-go-lib/logger"
	"github.com/stakkato95/twitter-service-graphql/config"
	"github.com/stakkato95/twitter-service-graphql/http/dto"
)

var tweetsService = config.AppConfig.TweetsService

type TweetRepo interface {
	CreateTweet(*dto.Tweet) (*dto.Tweet, error)
	GetTweets(int) ([]dto.Tweet, error)
}

type defaultTweetRepo struct {
}

func NewTweetRepo() TweetRepo {
	return &defaultTweetRepo{}
}

func (r *defaultTweetRepo) CreateTweet(tweet *dto.Tweet) (*dto.Tweet, error) {
	jsonData, err := json.Marshal(tweet)
	if err != nil {
		logger.Fatal("can not encode tweet: " + err.Error())
		return nil, err
	}

	response, err := http.DefaultClient.Post("http://localhost/tweets/debug/tweets", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		logger.Fatal("POST request to tweets service failed: " + err.Error())
		return nil, err
	}
	defer response.Body.Close()

	responseDto := dto.ResponseDto{}
	if err := json.NewDecoder(response.Body).Decode(&responseDto); err != nil {
		logger.Fatal("can not decode response from tweets service: " + err.Error())
		return nil, err
	}

	if responseDto.Error != "" {
		return nil, errors.New(responseDto.Error)
	}

	jsonData, err = json.Marshal(responseDto.Data)
	if err != nil {
		return nil, errors.New("can not marshal tweet data: " + err.Error())
	}

	createdTweet := dto.Tweet{}
	if err := json.NewDecoder(bytes.NewBuffer(jsonData)).Decode(&createdTweet); err != nil {
		logger.Fatal("can not decode tweet from data: " + err.Error())
		return nil, err
	}

	return &createdTweet, nil
}

func (r *defaultTweetRepo) GetTweets(userId int) ([]dto.Tweet, error) {
	return nil, nil
}

func mapToTweetDto(tweetMap map[string]interface{}) *dto.Tweet {
	return &dto.Tweet{
		Id:     int(tweetMap["id"].(float64)),
		UserId: int(tweetMap["userId"].(float64)),
		Text:   tweetMap["text"].(string),
	}
}
