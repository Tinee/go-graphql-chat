package graphql

import (
	context "context"
	"time"

	"github.com/Tinee/go-graphql-chat/domain"
)

func (r *mutationResolver) Register(ctx context.Context, input NewUser) (Viewer, error) {
	u, err := r.u.Create(domain.User{
		Password: input.Password,
		Username: input.Username,
	})
	if err != nil {
		return Viewer{}, err
	}
	tkn := r.claimJWT(u)

	return Viewer{ID: u.ID, Token: tkn, Username: u.Username}, nil
}

func (r *mutationResolver) Login(ctx context.Context, input LoginInput) (Viewer, error) {
	u, err := r.u.Authenticate(input.Username, input.Password)
	if err != nil {
		return Viewer{}, err
	}
	tkn := r.claimJWT(*u)

	return Viewer{ID: u.ID, Token: tkn, Username: u.Username}, nil
}

func (r *mutationResolver) PostMessage(ctx context.Context, text, username, roomID string) (Message, error) {
	// u, err := r.u.Authenticate(input.Username, input.Password)
	// if err != nil {
	// 	return Viewer{}, err
	// }
	// tkn := r.claimJWT(*u)

	return Message{
		CreatedAt: time.Now(),
		ID:        "",
		Text:      "",
	}, nil
}
