package domain

import (
	"errors"

	"github.com/stakkato95/twitter-service-graphql/http/dto"
)

type UserRepo interface {
	Create(*dto.UserDto) (*dto.NewUserDto, error)
	Authenticate(*dto.UserDto) (*dto.TokenDto, error)
}

type defaultUserRepo struct {
}

func NewUserRepo() UserRepo {
	return &defaultUserRepo{}
}

func (r *defaultUserRepo) Create(user *dto.UserDto) (*dto.NewUserDto, error) {
	return &dto.NewUserDto{
		User: *user,
		Token: dto.TokenDto{
			Token: "newtoken",
		},
	}, errors.New("no error")
}

func (r *defaultUserRepo) Authenticate(*dto.UserDto) (*dto.TokenDto, error) {
	return &dto.TokenDto{Token: "newtoken"}, errors.New("no error")
}
