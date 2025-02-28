package adapter

import (
	"fmt"
	"log"
	"sync"

	"meme-crawler/adapter/core"
)

// PostData is the data of a post
type PostData struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

// Child represents a post in Reddit's response
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

// RedditAdapter creates a new Reddit adapter
func RedditAdapter(maxResults int, userAgent string) *core.Adapter[RedditResponse] {
	return core.NewAdapter[RedditResponse](maxResults, userAgent)
}

// RedditMemeClient integrates Reddit
// use *sync.WaitGroup to wait for the integration to complete
// wg means wait group pointer
func RedditMemeClient(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("🔥 Start to integrate Reddit...")

	adapter := RedditAdapter(20, "MyRedditBot/1.0")
	result, err := adapter.Integrate("https://www.reddit.com/r/ProgrammerHumor/new.json", "Reddit", "")
	if err != nil {
		log.Printf("Error integrating Reddit: %v", err)
		return
	}
	// print the title and url of the post
	for _, post := range result.Data.Children {
		fmt.Printf("Title: %s\nURL: %s\n", post.Data.Title, post.Data.URL)
	}

	fmt.Println("🔥 Reddit integration completed, found", len(result.Data.Children), "posts")
}
