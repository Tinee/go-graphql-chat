package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Tinee/go-graphql-chat/middleware"

	"github.com/go-chi/chi"
)

func TestExtractTokenToContext(t *testing.T) {
	mux := chi.NewMux()
	mux.Use(middleware.ExtractTokenToContext)
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tkn := middleware.GetToken(r.Context())
		w.Write([]byte(tkn))
	})

	const tkn = "thisIsAFakeToken"
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r.Header.Set("Authorization", "Bearer "+tkn)

	mux.ServeHTTP(w, r)

	if w.Body.String() != tkn {
		t.Errorf("Expected ( %v ) to be ( %v )", w.Body.String(), tkn)
	}
}
