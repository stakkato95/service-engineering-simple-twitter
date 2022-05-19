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
)

func main() {
	router := chi.NewRouter()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	logger.Info("graphql service listening on port " + config.AppConfig.ServerPort)
	logger.Fatal("can not run server " + http.ListenAndServe(config.AppConfig.ServerPort, router).Error())
}
