package core

import (
	"fmt"
	"sync"
)

// ClientInterface interface defines the common behavior for all meme clients
type ClientInterface[T any] interface {
	Fetch(url string, websiteName string, token string, wg *sync.WaitGroup)
}

// Client creates a new Client
type Client[T any] struct {
	adapter *Adapter[T]
	Channel chan *T
}

// NewClient creates a new Client instance
func NewClient[T any](adapter *Adapter[T]) *Client[T] {
	return &Client[T]{
		adapter: adapter,
		Channel: make(chan *T),
	}
}

// Fetch implements ClientInterface
func (c *Client[T]) Fetch(
	url string,
	websiteName string,
	token string,
	wg *sync.WaitGroup,
) {
	if wg != nil {
		defer wg.Done()
	}

	// Get results from adapter
	results, err := c.adapter.Integrate(url, websiteName, token)
	if err != nil {
		fmt.Println("Error fetching results:", err)
		return
	}

	// Send results to channel
	c.Channel <- results
}
