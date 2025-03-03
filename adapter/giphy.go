package adapter

import (
	"fmt"
	"log"
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

// GiphyMemeClient integrates Giphy
func GiphyMemeClient(target string) ([]string, error) {

	if target == "" {
		target = "memes"
	}
	apiKey := os.Getenv("GIPHY_API_KEY")
	endpoint := "https://api.giphy.com/v1/gifs/search?api_key=" + apiKey + "&q=" + target + "&limit=25"

	fmt.Println("🔥 Start to integrate Giphy...")
	adapter := GiphyAdapter(20, "MyGiphyBot/1.0")
	result, err := adapter.Integrate(endpoint, "Giphy", "")
	if err != nil {
		log.Printf("Error integrating Giphy: %v", err)
		return nil, err
	}

	var urls []string
	for _, post := range result.Data {
		urls = append(urls, post.Images.Original.MP4)
	}

	fmt.Println("🔥 Giphy integration completed, found", len(result.Data), "posts")
	return urls, nil
}
