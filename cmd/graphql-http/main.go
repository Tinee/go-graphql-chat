package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"

	"github.com/Tinee/go-graphql-chat/graphql"
	"github.com/Tinee/go-graphql-chat/inmemory"
	"github.com/Tinee/go-graphql-chat/middleware"
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
)

func main() {
	var (
		log    = logrus.New()
		mux    = chi.NewMux()
		inmem  = inmemory.NewClient()
		ur     = inmem.UserRepository()
		ms     = inmem.MessageRepository()
		p      = inmem.ProfileRepository()
		port   = getEnvOrDefault("APP_PORT", "8080")
		secret = getEnvOrDefault("APP_SECRET", "localSecret")
	)
	log.Out = os.Stdout
	log.SetFormatter(&logrus.JSONFormatter{})

	err := inmem.FillWithMockData("./inmemory/mock_data.json")
	if err != nil {
		fmt.Printf("Couldn't load the mock data: %v \n", err)
	}
	mux.Use(
		cors.AllowAll().Handler,
		chiMiddleware.RequestID,
		chiMiddleware.Recoverer,
		middleware.TokenLifter,
	)

	mux.Handle(
		"/graphql",
		graphql.NewGraphQLHandlerFunc(ur, ms, p, log, secret),
	)

	fmt.Printf("Now listening on port :%v\n", port)
	http.ListenAndServe(":"+port, mux)

}

func getEnvOrDefault(key, d string) string {
	e := os.Getenv(key)
	if e == "" {
		return d
	}
	return e
}
