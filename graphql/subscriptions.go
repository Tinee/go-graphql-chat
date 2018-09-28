package graphql

import (
	"context"
)

func (r subscriptionResolver) MessageAdded(ctx context.Context, id string) (<-chan Message, error) {
	c := r.ls.addListener(ctx, id)

	return c, nil
}
