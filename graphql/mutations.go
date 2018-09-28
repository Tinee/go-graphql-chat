package graphql

import (
	context "context"

	"github.com/Tinee/go-graphql-chat/domain"
)

func (r *mutationResolver) Register(ctx context.Context, input NewUser) (User, error) {
	u, err := r.u.Create(domain.User{
		Password: input.Password,
		Username: input.Username,
	})
	if err != nil {
		return User{}, err
	}
	tkn := r.claimJWT(u)

	return User{ID: u.ID, Token: tkn, Username: u.Username}, nil
}

func (r *mutationResolver) Login(ctx context.Context, input LoginInput) (User, error) {
	u, err := r.u.Authenticate(input.Username, input.Password)
	if err != nil {
		return User{}, err
	}
	tkn := r.claimJWT(*u)

	return User{ID: u.ID, Token: tkn, Username: u.Username}, nil
}
