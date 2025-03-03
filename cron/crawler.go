package cron

import (
	"fmt"
	"log"
	"sync"
	"time"

	adapter "meme-crawler/adapter"

	"github.com/robfig/cron/v3"
)

// InitCrawler initializes the crawler
func InitCrawler() {
	cron := cron.New(cron.WithLocation(time.UTC))

	// Run the crawler every Sunday at 24:00
	if _, err := cron.AddFunc("0 0 * * 0", func() {
		fmt.Println("🔥 Start to crawl Reddit and Imgur...")

		var wg sync.WaitGroup
		wg.Add(2)
		go adapter.RedditMemeClient(&wg)
		go adapter.ImgurMemeClient(&wg)
		wg.Wait()
	}); err != nil {
		log.Printf("Failed to add cron job: %v", err)
		return
	}

	cron.Start()

	// Keep the program running
	select {}
}
