package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/Tinee/go-graphql-chat/graphql"
	"github.com/Tinee/go-graphql-chat/inmemory"
	"github.com/Tinee/go-graphql-chat/middleware"
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
)

func main() {
	var (
		mux    = chi.NewMux()
		inmem  = inmemory.NewClient()
		ur     = inmem.UserRepository()
		ms     = inmem.MessageRepository()
		p      = inmem.ProfileRepository()
		port   = getEnvOrDefault("APP_PORT", "8080")
		secret = getEnvOrDefault("APP_SECRET", "localSecret")
	)
	err := inmem.FillWithMockData()
	if err != nil {
		fmt.Printf("Couldn't load the mock data: %v", err)
	}
	mux.Use(
		cors.AllowAll().Handler,
		chiMiddleware.RequestID,
		chiMiddleware.Recoverer,
		middleware.TokenLifter,
	)

	r := graphql.New(ur, ms, p, secret)
	mux.Handle("/graphql", handler.GraphQL(r,
		handler.WebsocketUpgrader(websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		})))

	http.ListenAndServe(":"+port, mux)
}

func getEnvOrDefault(key, d string) string {
	e := os.Getenv(key)
	if e == "" {
		return d
	}
	return e
}
