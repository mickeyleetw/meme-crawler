package cmd

import (
	cron "meme-crawler/cron"

	"github.com/spf13/cobra"
)

// RunServerCmd is the command to run the server
var (
	RunCrawlerCmd = &cobra.Command{
		Use:   "run_crawler",
		Short: "Crawler commands",
		Long:  "Crawler commands",
		Run: func(cmd *cobra.Command, args []string) {
			cron.InitCrawler()
		},
	}
)
