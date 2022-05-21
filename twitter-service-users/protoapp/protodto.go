package protoapp

import (
	"github.com/stakkato95/twitter-service-users/domain"
	pb "github.com/stakkato95/twitter-service-users/proto"
)

func ToEntity(u *pb.User) domain.User {
	return domain.User{
		Id:       u.Id,
		Username: u.Username,
		Password: u.Password,
	}
}

func ToDto(u *domain.User, token string) *pb.NewUser {
	return &pb.NewUser{
		User: &pb.User{
			Id:       u.Id,
			Username: u.Username,
			Password: u.Password,
		},
		Token: &pb.Token{
			Token: token,
		},
	}
}
