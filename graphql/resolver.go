//go:generate gorunpkg github.com/99designs/gqlgen
package graphql

import (
	"sync"

	graphql "github.com/99designs/gqlgen/graphql"
	"github.com/Tinee/go-graphql-chat/domain"
)

type Resolver struct {
	u      domain.UserRepository
	ms     domain.MessageRepository
	p      domain.ProfileRepository
	ls     *listenerPool
	secret string
}

func New(
	u domain.UserRepository,
	ms domain.MessageRepository,
	p domain.ProfileRepository,
	secret string) graphql.ExecutableSchema {

	return NewExecutableSchema(Config{
		Resolvers: &Resolver{
			u:  u,
			ms: ms,
			p:  p,
			ls: &listenerPool{
				mtx: sync.Mutex{},
				ls:  make(map[string]*listener),
			},
			secret: secret,
		}})
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) Subscription() SubscriptionResolver {
	return &subscriptionResolver{r}
}

type subscriptionResolver struct{ *Resolver }

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }
