package graphql

import (
	"context"
	time "time"

	graphql "github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/handler"
	"github.com/Sirupsen/logrus"
	"github.com/go-chi/chi/middleware"
)

func OnErrorLogger(log *logrus.Logger) handler.Option {
	return handler.ResolverMiddleware(func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
		defer func(t time.Time) {
			if err != nil {
				log.WithFields(logrus.Fields{
					"requestId": middleware.GetReqID(ctx),
					"error":     err,
					"took":      time.Since(t),
				}).Error("Error happend when calling a resolver")
			}

		}(time.Now())
		return next(ctx)
	})
}

func RequestLogger(log *logrus.Logger) handler.Option {
	return handler.RequestMiddleware(func(ctx context.Context, next func(ctx context.Context) []byte) (res []byte) {
		defer func(t time.Time) {
			log.WithFields(logrus.Fields{
				"requestId": middleware.GetReqID(ctx),
				"took":      time.Since(t),
				"data":      string(res),
			}).Info("Successfully made an request.")
		}(time.Now())
		return next(ctx)
	})
}
