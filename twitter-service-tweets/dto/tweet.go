package dto

type TweetDto struct {
	Id     int    `json:"id,omitempty"`
	UserId int    `json:"userId" binding:"required"`
	Text   string `json:"text" binding:"required"`
}
