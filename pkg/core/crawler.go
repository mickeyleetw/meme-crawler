package core

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

// APICrawler is a crawler that crawls a website using an API
type APICrawler[T any] struct {
	maxResults int
	userAgent  string
}

// NewAPICrawler creates a new APICrawler
func NewAPICrawler[T any](maxResults int, userAgent string) *APICrawler[T] {
	return &APICrawler[T]{
		maxResults: maxResults,
		userAgent:  userAgent,
	}
}

func (c *APICrawler[T]) Crawl(wg *sync.WaitGroup, url string, websiteName string) (*T, error) {
	defer wg.Done()
	fmt.Println("🔥 Start crawling", websiteName, "using API...")

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", c.userAgent)

	resp, err := HTTPClientPool().Do(req)
	if err != nil {
		log.Println("❌", websiteName, "API request failed:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var result T
	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("❌", websiteName, "JSON parse failed:", err)
		return nil, err
	}

	return &result, nil
}
