package inmemory

import (
	"encoding/json"
	"math/rand"
	"os"
	"path/filepath"
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

type mockData struct {
	Profiles []domain.Profile `json:"profiles"`
	Users    []domain.User    `json:"users"`
}

// FillWithMockData this is when I realized that it sucks to not having a database in development.
func (c *Client) FillWithMockData(path string) error {
	abs, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	f, err := os.Open(abs)
	if err != nil {
		return err
	}

	var data mockData
	err = json.NewDecoder(f).Decode(&data)
	if err != nil {
		return err
	}

	for _, p := range data.Profiles {
		c.p.profiles = append(c.p.profiles, p)
	}
	for _, u := range data.Users {
		c.u.users = append(c.u.users, u)
	}

	return nil
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
