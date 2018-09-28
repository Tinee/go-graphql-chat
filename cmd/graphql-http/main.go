package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/Tinee/go-graphql-chat/graphql"
	"github.com/Tinee/go-graphql-chat/inmemory"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

func main() {
	mux := chi.NewMux()
	inmem := inmemory.NewClient()
	ur := inmem.UserRepository()
	var (
		port   = getEnvOrDefault("APP_PORT", "8080")
		secret = getEnvOrDefault("APP_SECRET", "localSecret")
	)

	mux.Use(
		cors.AllowAll().Handler,
	)
	res := graphql.NewResolver(ur, secret)

	mux.Handle("/query", handler.GraphQL(graphql.NewExecutableSchema(graphql.Config{Resolvers: res})))
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", "8080")
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func getEnvOrDefault(key, d string) string {
	e := os.Getenv(key)
	if e == "" {
		return d
	}
	return e
}
