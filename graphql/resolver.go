//go:generate gorunpkg github.com/99designs/gqlgen
package graphql

import (
	"net/http"
	"sync"

	"github.com/99designs/gqlgen/handler"
	"github.com/Sirupsen/logrus"
	"github.com/Tinee/go-graphql-chat/domain"
	"github.com/gorilla/websocket"
)

type Resolver struct {
	u      domain.UserRepository
	ms     domain.MessageRepository
	p      domain.ProfileRepository
	ls     *listenerPool
	log    *logrus.Logger
	secret string
}

func NewGraphQLHandlerFunc(
	u domain.UserRepository,
	ms domain.MessageRepository,
	p domain.ProfileRepository,
	log *logrus.Logger,
	secret string,
) http.HandlerFunc {

	schema := NewExecutableSchema(Config{
		Resolvers: &Resolver{
			u:   u,
			ms:  ms,
			p:   p,
			log: log,
			ls: &listenerPool{
				mtx: sync.Mutex{},
				ls:  make(map[string]*listener),
			},
			secret: secret,
		}})

	return handler.GraphQL(schema,
		RequestLogger(log),
		OnErrorLogger(log),
		handler.WebsocketUpgrader(websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}))
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) Viewer() ViewerResolver {
	return &viewerResolver{r}
}

func (r *Resolver) Subscription() SubscriptionResolver {
	return &subscriptionResolver{r}
}

type subscriptionResolver struct{ *Resolver }

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }

type viewerResolver struct{ *Resolver }
