package domain

import (
	"fmt"
	"reflect"

	"github.com/stakkato95/service-engineering-go-lib/logger"
)

type SubscriptionRepo interface {
	AddSubscription(int, Subscription)
}

type simpleSubscriptionRepo struct {
	repo DbRepo
}

func NewSubscriptionRepo(repo DbRepo) SubscriptionRepo {
	return &simpleSubscriptionRepo{repo}
}

func (r *simpleSubscriptionRepo) AddSubscription(from int, subscription Subscription) {
	user := User{}
	r.repo.GetDb().Where("id = ?", from).Find(&user)
	isNotInitialized := reflect.DeepEqual(user, User{})

	user.Subscriptions = append(user.Subscriptions, subscription)

	if isNotInitialized {
		r.repo.GetDb().Create(&user)
	} else {
		r.repo.GetDb().Model(&user).Updates(&user)
	}

	logger.Info(fmt.Sprintf("added subscription with id: %d", subscription.Id))
}
