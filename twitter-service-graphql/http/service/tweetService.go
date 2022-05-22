package service

import (
	"github.com/stakkato95/twitter-service-graphql/graph/model"
	"github.com/stakkato95/twitter-service-graphql/http/domain"
	"github.com/stakkato95/twitter-service-graphql/http/dto"
)

type TweetService interface {
	CreateTweet(model.NewTweet) (*model.Tweet, error)
}

type defaultTweetService struct {
	repo domain.TweetRepo
}

func NewTweetService(repo domain.TweetRepo) TweetService {
	return &defaultTweetService{repo}
}

func (s *defaultTweetService) CreateTweet(tweet model.NewTweet) (*model.Tweet, error) {
	tweetDto := dto.TweetToDto(tweet)
	createdTweet, err := s.repo.CreateTweet(tweetDto)
	if err != nil {
		return nil, err
	}

	return dto.TweetDtoToGraphql(*createdTweet), nil
}
