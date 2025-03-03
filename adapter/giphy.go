package adapter

import (
	"fmt"
	"os"

	"meme-crawler/adapter/core"
)

type originalImage struct {
	MP4 string `json:"mp4"`
}

// GiphyImageData represents the image data structure from Giphy API
type GiphyImageData struct {
	Original originalImage `json:"original"`
}

// GiphyData is the data of a Giphy API
type GiphyData struct {
	Images GiphyImageData `json:"images"`
}

// GiphyResponse is the response of a Giphy API
type GiphyResponse struct {
	Data []GiphyData `json:"data"`
}

// GiphyAdapter creates a new Giphy adapter
func GiphyAdapter(maxResults int, userAgent string) *core.Adapter[GiphyResponse] {
	return core.NewAdapter[GiphyResponse](maxResults, userAgent)
}

// GiphyClient creates a new Giphy client
func GiphyClient(maxResults int, userAgent string) *core.Client[GiphyResponse] {
	return core.NewClient(GiphyAdapter(maxResults, userAgent))
}

// GiphyMemeClient integrates Giphy
func GiphyMemeClient(target string) []string {
	if target == "" {
		target = "memes"
	}
	fmt.Println("🔥 Start to integrate Giphy...")
	apiKey := os.Getenv("GIPHY_API_KEY")
	endpoint := "https://api.giphy.com/v1/gifs/search?api_key=" + apiKey + "&q=" + target + "&limit=25"

	client := GiphyClient(20, "MyGiphyBot/1.0")
	client.Fetch(endpoint, "Giphy", "", nil)

	var urls []string
	if result := <-client.Channel; result != nil {
		urls = make([]string, 0, len(result.Data))
		for _, post := range result.Data {
			urls = append(urls, post.Images.Original.MP4)
		}
		fmt.Println("🔥 Giphy integration completed, found", len(result.Data), "posts")
	}

	// Close the channel after processing
	close(client.Channel)

	return urls
}
