package chatbot

import (
	"fmt"
	"log"
	"net/http"
	"os"

	adapter "meme-crawler/adapter"

	"github.com/line/line-bot-sdk-go/v8/linebot"
)

var (
	lineBot *linebot.Client
	err     error
)

// InitLineBot initializes the Line bot
func InitLineBot() {
	channelSecret := os.Getenv("LINE_CHANNEL_SECRET")
	channelToken := os.Getenv("LINE_CHANNEL_TOKEN")
	lineBot, err = linebot.New(channelSecret, channelToken)
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/callback", callbackHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("🔥 Line bot initialized")
		w.WriteHeader(http.StatusOK)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := lineBot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			http.Error(w, "Invalid signature", http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			userID := event.Source.UserID

			if message, ok := event.Message.(*linebot.TextMessage); ok {
				giphyGIFs := adapter.GiphyMemeClient(message.Text)
				if err != nil {
					log.Println("Error fetching Giphy:", err)
					return
				}
				if _, err := lineBot.PushMessage(userID, linebot.NewTextMessage("GIFs for you!")).Do(); err != nil {
					log.Println("Error sending message:", err)
				}
				for _, giphyGIF := range giphyGIFs {
					if _, err := lineBot.PushMessage(userID, linebot.NewVideoMessage(giphyGIF, giphyGIF)).Do(); err != nil {
						log.Println("Error sending message:", err)
					}
				}
			} else {
				if _, err := lineBot.PushMessage(userID, linebot.NewTextMessage("Sorry, I can't process this message!")).Do(); err != nil {
					log.Println("Error sending error message:", err)
				}
			}
		}
	}
}
