package adapter

import (
	"fmt"
	"log"
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

// ImgurMemeClient integrates Imgur
// use *sync.WaitGroup to wait for the integration to complete
// wg means wait group pointer
func ImgurMemeClient(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("🔥 Start to integrate Imgur...")

	adapter := ImgurAdapter(20, "MyImgurBot/1.0")
	result, err := adapter.Integrate("https://api.imgur.com/3/gallery/search/time/0?q=memes", "Imgur", "22e14abbb66ad685470ca3201dfb1b8b9613defe")
	if err != nil {
		log.Printf("Error integrating Imgur: %v", err)
		return
	}
	// print the title and url of the post
	for _, post := range result.Data {
		for _, image := range post.Images {
			fmt.Printf("Id: %s\nLink: %s\n", image.ID, image.Link)
		}
	}

	fmt.Println("🔥 Imgur integration completed, found", len(result.Data), "posts")
}
