package domain

import "github.com/stakkato95/twitter-service-tweets/dto"

type Tweet struct {
	// gorm.Model
	Id     int
	UserId int
	Text   string
}

func ToEntity(dto *dto.TweetDto) *Tweet {
	return &Tweet{
		Id:     dto.Id,
		UserId: dto.UserId,
		Text:   dto.Text,
	}
}

func ToDto(entity *Tweet) *dto.TweetDto {
	return &dto.TweetDto{
		Id:     entity.Id,
		UserId: entity.UserId,
		Text:   entity.Text,
	}
}
