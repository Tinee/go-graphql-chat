package graphql_test

import (
	"net/http/httptest"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/Sirupsen/logrus"
	"github.com/Tinee/go-graphql-chat/graphql"
	"github.com/Tinee/go-graphql-chat/inmemory"
)

func Test_graphql_mutationResolver(t *testing.T) {
	inmem := inmemory.NewClient()
	err := inmem.FillWithMockData("../inmemory/mock_data.json")

	if err != nil {
		t.Fatal("Error: We need to have mock data inserted.")
	}

	srv := httptest.NewServer(graphql.NewGraphQLHandlerFunc(
		inmem.UserRepository(),
		inmem.MessageRepository(),
		inmem.ProfileRepository(),
		logrus.New(),
		"localSecret",
	))
	c := client.New(srv.URL)

	t.Run("Mutation Register", func(t *testing.T) {
		var resp struct {
			Register struct {
				ID       string
				Username string
			}
		}
		c.MustPost(`mutation { register(input: { username: "Marcus", password: "admin" }) { id, username } }`, &resp)

		if resp.Register.ID == "" {
			t.Error("Didn't expect this to be empty.")
		}
		if resp.Register.Username != "Marcus" {
			t.Errorf("Expected username to be (%v) but got (%v)", "Marcus", resp.Register.Username)
		}
	})

	t.Run("Mutation Login", func(t *testing.T) {
		var resp struct {
			Login struct {
				ID       string
				Username string
				Token    string
			}
		}
		c.MustPost(`mutation { login(input: { username: "tine", password: "test1" }) {  username } }`, &resp)

		if resp.Login.Username != "tine" {
			t.Errorf("Expected (%v) but got (%v)", "tine", resp.Login.Username)
		}
	})

	t.Run("Mutation PostMessage", func(t *testing.T) {
		var resp struct {
			PostMessage struct {
				SenderID string
			}
		}
		c.MustPost(`mutation { postMessage(input:{text:"Foo", senderId:"2", receiverId:"1"}) { senderId } }`, &resp)

		if resp.PostMessage.SenderID != "2" {
			t.Errorf("Expected (%v) but got (%v)", "2", resp.PostMessage.SenderID)
		}
	})

	t.Run("Mutation PostProfile", func(t *testing.T) {
		var resp struct {
			PostProfile struct {
				UserID string
			}
		}
		c.MustPost(`mutation { postProfile(input:{ userId:"Foo", firstName:"Foo", lastName:"Bar", age: 25 }) { userId } }`, &resp)

		if resp.PostProfile.UserID != "Foo" {
			t.Errorf("Expected (%v) but got (%v)", "Foo", resp.PostProfile.UserID)
		}
	})

}
