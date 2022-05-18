package dto

import "github.com/stakkato95/twitter-service-users/domain"

type UserDto struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *UserDto) ToEntity() domain.User {
	return domain.User{
		Id:       u.Id,
		Username: u.Username,
		Password: u.Password,
	}
}

func ToDto(u domain.User) UserDto {
	return UserDto{
		Id:       u.Id,
		Username: u.Username,
		Password: u.Password,
	}
}
