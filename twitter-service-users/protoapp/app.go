package protoapp

import (
	"net"

	"github.com/stakkato95/service-engineering-go-lib/logger"
	grpc "google.golang.org/grpc"

	"github.com/stakkato95/twitter-service-users/config"
	pb "github.com/stakkato95/twitter-service-users/proto"
)

func Start() {
	lis, err := net.Listen("tcp", config.AppConfig.GrpcPort)
	if err != nil {
		logger.Fatal("can not listen on grpc server port: " + err.Error())
	}

	server := grpc.NewServer()
	pb.RegisterUsersServiceServer(server, &defaultUsersServiceServer{})
	logger.Info("users grpc service listening on port " + config.AppConfig.GrpcPort)
	logger.Fatal("can not run grpc server: " + server.Serve(lis).Error())
}
