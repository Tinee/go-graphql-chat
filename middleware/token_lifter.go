package middleware

import (
	"context"
	"net/http"
	"strings"
)

var tokenCtxKey = &contextKey{"token"}

type contextKey struct {
	name string
}

// TokenLifter check if we have a Authorization Header with prefix Bearer
// if that's the case it lifts it into the context.
func TokenLifter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		token = strings.TrimPrefix(token, "Bearer ")

		ctx := context.WithValue(r.Context(), tokenCtxKey, token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetToken returns a token from the given context if one is present.
// Returns the empty string if a token couldn't be found.
func GetToken(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if reqID, ok := ctx.Value(tokenCtxKey).(string); ok {
		return reqID
	}
	return ""
}
