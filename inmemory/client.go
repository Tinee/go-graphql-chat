package inmemory

import (
	"math/rand"
	"sync"
	"time"

	"github.com/Tinee/go-graphql-chat/domain"
)

type Client struct {
	u *userInMemory
}

func NewClient() *Client {
	return &Client{
		u: &userInMemory{
			mtx: &sync.Mutex{},
		},
	}
}

func (c *Client) UserRepository() domain.UserRepository {
	return c.u
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
