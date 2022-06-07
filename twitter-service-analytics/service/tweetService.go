package service

import "github.com/stakkato95/twitter-service-analytics/domain"

type TweetService interface {
	GetTweetCount() int
}

type simpleTweetService struct {
	repo domain.TweetProcessor
}

func NewTweetService(repo domain.TweetProcessor) TweetService {
	return &simpleTweetService{repo}
}

func (s *simpleTweetService) GetTweetCount() int {
	return s.repo.GetTweetCount()
}
