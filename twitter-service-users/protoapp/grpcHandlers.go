package protoapp

import (
	"context"

	pb "github.com/stakkato95/twitter-service-users/proto"
)

type defaultUsersServiceServer struct {
	pb.UnimplementedUsersServiceServer
}

func (defaultUsersServiceServer) CreateUser(ctx context.Context, user *pb.User) (*pb.NewUser, error) {
	u := &pb.NewUser{
		User: &pb.User{
			Id:       1,
			Username: "user100500",
			Password: "pass",
		},
		Token: &pb.Token{Token: "tokenn"},
	}
	return u, nil
}

func (defaultUsersServiceServer) AuthUser(ctx context.Context, user *pb.User) (*pb.Token, error) {
	return &pb.Token{Token: "token100500"}, nil
}
