package dto

type SubscriptionDto struct {
	From int `json:"from" binding:"required"`
	To   int `json:"to" binding:"required"`
}
