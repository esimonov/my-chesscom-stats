package client

import (
	"net/http"
	"time"
)

// Client for communication with chess.com API servers.
type Client struct {
	Username       string
	internalClient *http.Client
}

// NewClient returns pointer to new client.
func NewClient(username string) *Client {
	client := &Client{
		Username:       username,
		internalClient: http.DefaultClient,
	}

	client.internalClient.Timeout = 10 * time.Second
	return client
}
