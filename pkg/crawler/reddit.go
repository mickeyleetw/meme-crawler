package crawler

import (
	"fmt"
	"log"
	"sync"

	"meme-crawler/pkg/core"
)

// PostData is the data of a post
type PostData struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

// Child is a child of a post
type Child struct {
	Data PostData `json:"data"`
}

// ResponseData is the data of a response
type ResponseData struct {
	Children []Child `json:"children"`
}

// RedditResponse is the response of a Reddit API
type RedditResponse struct {
	Data ResponseData `json:"data"`
}

// RedditCrawler creates a new Reddit crawler
func RedditCrawler(maxResults int, userAgent string) *core.APICrawler[RedditResponse] {
	return core.NewAPICrawler[RedditResponse](maxResults, userAgent)
}

// CrawlReddit crawls Reddit
func CrawlReddit() {
	var wg sync.WaitGroup
	wg.Add(1)

	crawler := RedditCrawler(20, "MyRedditBot/1.0")
	result, err := crawler.Crawl(&wg, "https://www.reddit.com/r/memes/top.json?limit=20", "Reddit")
	if err != nil {
		log.Fatal(err)
	}

	for _, post := range result.Data.Children {
		fmt.Println(post.Data.Title)
	}
}
