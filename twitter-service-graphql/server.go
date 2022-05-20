package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/stakkato95/service-engineering-go-lib/logger"
	"github.com/stakkato95/twitter-service-graphql/config"
	"github.com/stakkato95/twitter-service-graphql/graph"
	"github.com/stakkato95/twitter-service-graphql/graph/generated"
	"github.com/stakkato95/twitter-service-graphql/http/domain"
	"github.com/stakkato95/twitter-service-graphql/http/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/stakkato95/twitter-service-graphql/proto"
)

func main() {
	//3 add grpc to users service
	//4 add grpc calls to users service (users in k8s + graphql on localhost)
	//5 add grpc calls to users service (users in k8s + graphql in k8s)

	repo := domain.NewUserRepo()
	service := service.NewUserService(repo)

	router := chi.NewRouter()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{Service: service},
	}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	//
	conn, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal("can not listen to users grpc server: " + err.Error())
	}
	defer conn.Close()

	client := pb.NewUsersServiceClient(conn)

	timeout := 10
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	newUser, err := client.CreateUser(ctx, &pb.User{Username: "u", Password: "p"})
	if err != nil {
		logger.Fatal("can not create user via users grpc interface: " + err.Error())
	}
	logger.Info(fmt.Sprintf("newUser: %#v", newUser))

	logger.Info("graphql service listening on port " + config.AppConfig.ServerPort)
	logger.Fatal("can not run server " + http.ListenAndServe(config.AppConfig.ServerPort, router).Error())
}
