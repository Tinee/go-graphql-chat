package graphql

import (
	context "context"

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

func (r *mutationResolver) PostMessage(ctx context.Context, input NewMessage) (Message, error) {
	m, err := r.ms.Create(domain.Message{
		ReceiverID: input.ReceiverID,
		SenderID:   input.SenderID,
		Text:       input.Text,
	})
	if err != nil {
		return Message{}, err
	}

	out := Message{
		CreatedAt: m.CreatedAt,
		ID:        m.ID,
		SenderID:  m.SenderID,
		Text:      m.Text,
	}
	r.ls.sendMessage(m.ReceiverID, out)
	return out, nil
}

func (r *mutationResolver) PostProfile(ctx context.Context, input NewProfile) (Profile, error) {
	p, err := r.p.Create(domain.Profile{
		UserID:    input.UserID,
		Age:       input.Age,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	})
	if err != nil {
		return Profile{}, err
	}
	out := Profile{
		ID:        p.ID,
		Age:       p.Age,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		UserID:    p.UserID,
	}
	return out, nil
}
