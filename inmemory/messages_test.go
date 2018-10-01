package inmemory_test

import (
	"testing"

	"github.com/Tinee/go-graphql-chat/inmemory"

	"github.com/Tinee/go-graphql-chat/domain"
)

func Test_messagesInMemory_Create(t *testing.T) {
	c := NewClient()
	repo := c.MessageRepository()

	m, err := repo.Create(domain.Message{
		ReceiverID: "Foo",
		SenderID:   "Bar",
		Text:       "FooBar",
	})

	if err != nil {
		t.Errorf("Expected not an error but got: %v", err)
	}

	_, err = repo.Find(m.ID)
	if err == inmemory.ErrProfileNotFound {
		t.Error("Expected to find the entity, but I didn't.")
	}
}
