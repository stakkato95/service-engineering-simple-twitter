package service

import (
	"github.com/stakkato95/twitter-service-tweets/domain"
	"github.com/stakkato95/twitter-service-tweets/dto"
)

type SubscriptionService interface {
	AddSubscription(dto.SubscriptionDto)
}

type simpleSubscriptionService struct {
	repo domain.SubscriptionRepo
}

func NewSubscriptionService(repo domain.SubscriptionRepo) SubscriptionService {
	return &simpleSubscriptionService{repo}
}

func (s *simpleSubscriptionService) AddSubscription(subscriptionDto dto.SubscriptionDto) {
	s.repo.AddSubscription(subscriptionDto.From, domain.Subscription{To: subscriptionDto.To})
}
