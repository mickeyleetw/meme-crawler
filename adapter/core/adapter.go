package core

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Adapter is a crawler that crawls a website using an API
type Adapter[T any] struct {
	maxResults int
	userAgent  string
}

// NewAdapter creates a new Adapter
func NewAdapter[T any](maxResults int, userAgent string) *Adapter[T] {
	return &Adapter[T]{
		maxResults: maxResults,
		userAgent:  userAgent,
	}
}

func (c *Adapter[T]) Integrate(url string, websiteName string, token string) (*T, error) {

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", c.userAgent)
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	resp, err := HTTPClientPool().Do(req)
	if err != nil {
		log.Println("❌", websiteName, "API request failed:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("❌", websiteName, "Failed to read response body:", err)
		return nil, err
	}

	var result T
	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("❌", websiteName, "JSON parse failed:", err)
		return nil, err
	}

	return &result, nil
}
