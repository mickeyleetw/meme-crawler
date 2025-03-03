package adapter

import (
	"fmt"
	"sync"

	"meme-crawler/adapter/core"
)

// ImageData is the data of a image
type ImageData struct {
	ID   string `json:"id"`
	Link string `json:"link"`
}

// ImgurData is the data of a Imgur API
type ImgurData struct {
	Images []ImageData `json:"images"`
}

// ImgurResponse is the response of a Imgur API
type ImgurResponse struct {
	Data []ImgurData `json:"data"`
}

// ImgurAdapter creates a new Imgur adapter
func ImgurAdapter(maxResults int, userAgent string) *core.Adapter[ImgurResponse] {
	return core.NewAdapter[ImgurResponse](maxResults, userAgent)
}

// ImgurClient creates a new Imgur client
func ImgurClient(maxResults int, userAgent string) *core.Client[ImgurResponse] {
	return core.NewClient(ImgurAdapter(maxResults, userAgent))
}

// ImgurMemeClient integrates Imgur
// use *sync.WaitGroup to wait for the integration to complete
// wg means wait group pointer
func ImgurMemeClient(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("🔥 Start to integrate Imgur...")

	client := ImgurClient(20, "MyImgurBot/1.0")

	// Call Fetch with all required parameters
	client.Fetch(
		"https://api.imgur.com/3/gallery/search/time/0?q=memes",
		"Imgur",
		"22e14abbb66ad685470ca3201dfb1b8b9613defe",
		wg,
	)

	// Process the ImgurResponse and convert to []ImageData
	if result := <-client.Channel; result != nil {
		images := make([]ImageData, 0)
		for _, post := range result.Data {
			images = append(images, post.Images...)
		}
		fmt.Println("🔥 Imgur integration completed, found", len(images), "posts")
	}

	// Close the channel after processing
	close(client.Channel)
}
