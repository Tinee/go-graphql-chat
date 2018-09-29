package inmemory

import (
	"math/rand"
	"sync"
	"time"

	"github.com/Tinee/go-graphql-chat/domain"
)

type Client struct {
	u  *userInMemory
	ms *messagesInMemory
	p  *profileInMemory
}

func NewClient() *Client {
	return &Client{
		u: &userInMemory{
			mtx: &sync.Mutex{},
		},
		ms: &messagesInMemory{
			mtx: &sync.Mutex{},
		},
		p: &profileInMemory{
			mtx: &sync.Mutex{},
		},
	}
}

func (c *Client) UserRepository() domain.UserRepository {
	return c.u
}

func (c *Client) MessageRepository() domain.MessageRepository {
	return c.ms
}

func (c *Client) ProfileRepository() domain.ProfileRepository {
	return c.p
}

func generateID() string {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 20)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
