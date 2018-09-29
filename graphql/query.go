package graphql

import (
	"context"

	"github.com/Tinee/go-graphql-chat/middleware"
)

func (r *queryResolver) Me(ctx context.Context) (Viewer, error) {
	tkn := middleware.GetToken(ctx)
	id, err := r.validateAndExtractId(tkn)
	if err != nil {
		return Viewer{}, err
	}

	u, err := r.u.Find(id)
	if err != nil {
		return Viewer{}, err
	}
	out := Viewer{
		ID:       u.ID,
		Token:    tkn,
		Username: u.Username,
	}
	return out, nil
}
