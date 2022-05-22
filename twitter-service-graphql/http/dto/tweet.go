package dto

import (
	"github.com/stakkato95/twitter-service-graphql/graph/model"
)

type Tweet struct {
	Id     int    `json:"id"`
	UserId int    `json:"userId"`
	Text   string `json:"text"`
}

func TweetToDto(tweet model.NewTweet) *Tweet {
	return &Tweet{
		UserId: tweet.UserID,
		Text:   tweet.Text,
	}
}

func TweetDtoToGraphql(dto Tweet) *model.Tweet {
	return &model.Tweet{
		ID:     dto.Id,
		UserID: dto.UserId,
		Text:   dto.Text,
	}
}
