package main

import (
	"log"
	cmd "meme-crawler/cmd"

	"github.com/spf13/cobra"
)

// rootCmd is the root command for the application and & stands for the new command object pointer
var rootCmd = &cobra.Command{
	Use:   "meme-crawler",
	Short: "Meme Crawler",
	Long:  "Meme Crawler",
}

func init() {
	rootCmd.AddCommand(cmd.RunLineBotCmd)
	rootCmd.AddCommand(cmd.RunCrawlerCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Failed to execute root command: %v", err)
	}
}
