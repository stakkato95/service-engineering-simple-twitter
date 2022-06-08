package service

import (
	"fmt"

	"github.com/stakkato95/service-engineering-go-lib/logger"
	"github.com/stakkato95/twitter-service-tweets/domain"
	"github.com/stakkato95/twitter-service-tweets/dto"
)

type TweetsService interface {
	AddTweet(dto.TweetDto) *dto.TweetDto
	GetAllTweets(int) []dto.TweetDto
}

type defaultTweetsService struct {
	repo domain.TweetsRepo
	sink domain.TweetsSink
}

func NewTweetsService(repo domain.TweetsRepo, sink domain.TweetsSink) TweetsService {
	return &defaultTweetsService{repo, sink}
}

func (s *defaultTweetsService) AddTweet(tweetDto dto.TweetDto) *dto.TweetDto {
	entity := domain.ToEntity(&tweetDto)
	createdTweet := s.repo.AddTweet(*entity)
	if err := s.sink.AddTweet(tweetDto); err != nil {
		logger.Fatal(fmt.Sprintf("can not add tweet to sink: %v", tweetDto))
	}
	return domain.ToDto(createdTweet)
}

func (s *defaultTweetsService) GetAllTweets(userId int) []dto.TweetDto {
	tweets := s.repo.GetAllTweets(userId)
	tweetsDto := make([]dto.TweetDto, len(tweets))

	for i, tweet := range tweets {
		tweetsDto[i] = *domain.ToDto(&tweet)
	}

	return tweetsDto
}
