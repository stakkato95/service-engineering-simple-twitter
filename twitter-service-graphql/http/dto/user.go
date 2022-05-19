package dto

import "github.com/stakkato95/twitter-service-graphql/graph/model"

type UserDto struct {
	Id       int64  `json:"id,omitempty"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
}

func ToDtoFromUser(u model.NewUser) UserDto {
	return UserDto{
		Username: u.Username,
		Password: u.Password,
	}
}

func ToDtoFromLogin(u model.Login) UserDto {
	return UserDto{
		Username: u.Username,
		Password: u.Password,
	}
}
