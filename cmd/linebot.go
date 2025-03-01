package cmd

import (
	chatbot "meme-crawler/chatbot"

	"github.com/spf13/cobra"
)

// RunServerCmd is the command to run the server
var (
	RunLineBotCmd = &cobra.Command{
		Use:   "run",
		Short: "LineBot commands",
		Long:  "LineBot commands",
		Run: func(cmd *cobra.Command, args []string) {
			chatbot.InitLineBot()
		},
	}
)
