package main

import (
	"fmt"
	adapter "meme-crawler/adapter"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go adapter.RedditMemeClient(&wg)
	go adapter.ImgurMemeClient(&wg)

	wg.Wait()

	fmt.Println("🚀 Meme crawler completed")
}
