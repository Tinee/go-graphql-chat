package inmemory_test

import (
	"github.com/Tinee/go-graphql-chat/inmemory"
)

type Client struct {
	*inmemory.Client
}

func NewClient() *Client {
	inner := inmemory.NewClient()
	return &Client{inner}
}

func (c *Client) Reset() {
	c.Client = inmemory.NewClient()
}

func (c *Client) FillWithMockData() {
	c.Client.FillWithMockData("mock_data.json")
}
