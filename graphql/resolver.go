//go:generate gorunpkg github.com/99designs/gqlgen
package graphql

import (
	context "context"
	"sync"

	graphql "github.com/99designs/gqlgen/graphql"
	"github.com/Tinee/go-graphql-chat/domain"
)

type Resolver struct {
	u      domain.UserRepository
	ms     domain.MessageRepository
	ls     *listenerPool
	secret string
}

func New(u domain.UserRepository, ms domain.MessageRepository, secret string) graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: &Resolver{
			u:  u,
			ms: ms,
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

func (r *queryResolver) Me(ctx context.Context) (Viewer, error) {
	panic("not implemented")
}
