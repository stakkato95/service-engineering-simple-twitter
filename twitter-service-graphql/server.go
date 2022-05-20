package main

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/stakkato95/service-engineering-go-lib/logger"
	"github.com/stakkato95/twitter-service-graphql/config"
	"github.com/stakkato95/twitter-service-graphql/graph"
	"github.com/stakkato95/twitter-service-graphql/graph/generated"
	"github.com/stakkato95/twitter-service-graphql/http/domain"
	"github.com/stakkato95/twitter-service-graphql/http/service"
)

func main() {
	//1 add http calls to users service (users in k8s + graphql on localhost)
	//2 add http calls to users service (users in k8s + graphql in k8s)

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

	logger.Info("graphql service listening on port " + config.AppConfig.ServerPort)
	logger.Fatal("can not run server " + http.ListenAndServe(config.AppConfig.ServerPort, router).Error())
}
