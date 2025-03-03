package adapter

import (
	"fmt"
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

// RedditClient creates a new RedditClient
func RedditClient(maxResults int, userAgent string) *core.Client[RedditResponse] {
	return core.NewClient(RedditAdapter(maxResults, userAgent))
}

// RedditMemeClient integrates Reddit
// use *sync.WaitGroup to wait for the integration to complete
// wg means wait group pointer
// resultChan is a channel to send the result posts
func RedditMemeClient(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("🔥 Start to integrate Reddit...")

	client := RedditClient(20, "MyRedditBot/1.0")

	// Call Fetch with all required parameters
	client.Fetch(
		"https://www.reddit.com/r/ProgrammerHumor/new.json",
		"Reddit",
		"",
		wg,
	)

	// Process the RedditResponse and convert to []PostData
	if result := <-client.Channel; result != nil {
		posts := make([]PostData, 0, len(result.Data.Children))
		for _, post := range result.Data.Children {
			posts = append(posts, post.Data)
		}
		fmt.Println("🔥 Reddit integration completed, found", len(posts), "posts")
	}

	// Close the channel after processing
	close(client.Channel)
}
