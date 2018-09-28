//go:generate gorunpkg github.com/99designs/gqlgen
package graphql

import (
	context "context"

	"github.com/Tinee/go-graphql-chat/domain"
)

type Resolver struct {
	u      domain.UserRepository
	secret string
}

func NewResolver(u domain.UserRepository, secret string) *Resolver {
	return &Resolver{u, secret}
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }

func (r *queryResolver) User(ctx context.Context) (*User, error) {
	panic("not implemented")
}
