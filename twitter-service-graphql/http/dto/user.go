package dto

import (
	"github.com/stakkato95/twitter-service-graphql/graph/model"

	pb "github.com/stakkato95/twitter-service-graphql/proto"
)

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

func NewUserToProto(u *UserDto) *pb.User {
	return &pb.User{
		Id:       u.Id,
		Username: u.Username,
		Password: u.Password,
	}
}

func NewUserToDto(newUser *pb.NewUser) *NewUserDto {
	return &NewUserDto{
		User: UserDto{
			Id:       newUser.User.Id,
			Username: newUser.User.Username,
			Password: newUser.User.Password,
		},
		Token: TokenDto{newUser.Token.Token},
	}
}

func TokenToDto(token *pb.Token) *TokenDto {
	return &TokenDto{Token: token.Token}
}
