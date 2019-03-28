package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
)

func main() {
	if len(os.Args) < 2 {
		panic(errors.New("Bot token string must be provided as first argument."))
	}

	bot, err := newBot(os.Args[1], BotUpdateTimeout)
	browser, err := newBrowser()
	if err != nil {
		panic(err.Error())
	}

	updates, err := bot.GetUpdates()
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if reflect.TypeOf(update.Message.Text).Kind() != reflect.String || update.Message.Text == "" {
			bot.ReplyToMessage(*update.Message, "This message type is not supported.")
			continue
		}

		if update.Message.Text == "/start" {
			bot.sendMessage(*update.Message.Chat, "Hey, can you send me any URL?")
			continue
		}

		urls, err := ParseUrlsFromText(update.Message.Text)
		if err != nil {
			bot.ReplyToMessage(*update.Message, err.Error())
			continue
		}

		for _, url := range urls {
			err := browser.OpenUrl(url)
			if err != nil {
				log.Print(err.Error())
				continue
			}

			bot.ReplyToMessage(*update.Message, fmt.Sprintf("%s opened on TV!", url))
		}
	}
}
