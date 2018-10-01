package graphql_test

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/99designs/gqlgen/client"
	"github.com/Sirupsen/logrus"
	"github.com/Tinee/go-graphql-chat/graphql"
	"github.com/Tinee/go-graphql-chat/inmemory"
)

func Test_subscriptionResolver_MessageAdded(t *testing.T) {
	inmem := inmemory.NewClient()

	srv := httptest.NewServer(graphql.NewGraphQLHandlerFunc(
		inmem.UserRepository(),
		inmem.MessageRepository(),
		inmem.ProfileRepository(),
		logrus.New(),
		"localSecret",
	))
	c := client.New(srv.URL)

	sub := c.Websocket(`subscription { messageAdded(id: "1"){ text senderId } }`)
	defer sub.Close()

	go func() {
		var resp interface{}
		time.Sleep(10 * time.Millisecond)
		err := c.Post(`mutation {
				a:postMessage(input:{text:"Foo1", senderId:"2", receiverId:"1"}) { senderId text }
				b:postMessage(input:{text:"Foo2", senderId:"2", receiverId:"1"}) { senderId text }
			}`, &resp)

		if err != nil {
			t.Fatalf("Expected not to get an error here, bailing: ( %v )", err)
		}
	}()

	var msg struct {
		resp struct {
			MessageAdded struct {
				Text     string
				SenderID string
			}
		}
		err error
	}

	msg.err = sub.Next(&msg.resp)
	if msg.err != nil {
		t.Errorf("Expected not to get an error here:: ( %v )", msg.err)
	}
	if msg.resp.MessageAdded.Text != "Foo1" {
		t.Errorf("Expected to get (%v) but got ( %v )", "Foo1", msg.resp.MessageAdded.Text)
	}
	if msg.resp.MessageAdded.SenderID != "2" {
		t.Errorf("Expected to get (%v) but got ( %v )", "Foo1", msg.resp.MessageAdded.SenderID)
	}

	msg.err = sub.Next(&msg.resp)
	if msg.err != nil {
		t.Errorf("Expected not to get an error here:: ( %v )", msg.err)
	}
	if msg.resp.MessageAdded.Text != "Foo2" {
		t.Errorf("Expected to get (%v) but got ( %v )", "Foo1", msg.resp.MessageAdded.Text)
	}
	if msg.resp.MessageAdded.SenderID != "2" {
		t.Errorf("Expected to get (%v) but got ( %v )", "Foo1", msg.resp.MessageAdded.SenderID)
	}
}
