package service

import (
	"github.com/stakkato95/twitter-service-tweets/domain"
	"github.com/stakkato95/twitter-service-tweets/dto"
)

type TweetsService interface {
	AddTweet(dto.TweetDto) (*dto.TweetDto, error)
	GetAllTweets(int) ([]dto.TweetDto, error)
}

type defaultTweetsService struct {
	repo domain.TweetsRepo
}

func NewTweetsService(repo domain.TweetsRepo) TweetsService {
	return &defaultTweetsService{repo}
}

func (s *defaultTweetsService) AddTweet(tweetDto dto.TweetDto) (*dto.TweetDto, error) {
	entity := dto.ToEntity(&tweetDto)

	createdTweet, err := s.repo.AddTweet(*entity)
	if err != nil {
		return nil, err
	}

	return dto.ToDto(createdTweet), nil
}

func (s *defaultTweetsService) GetAllTweets(userId int) ([]dto.TweetDto, error) {
	tweets, err := s.repo.GetAllTweets(userId)
	if err != nil {
		return nil, err
	}

	tweetsDto := make([]dto.TweetDto, len(tweets))

	for i, tweet := range tweets {
		tweetsDto[i] = *dto.ToDto(&tweet)
	}

	return tweetsDto, nil
}
