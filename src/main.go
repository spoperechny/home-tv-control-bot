package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
)

func main() {
	botToken, err := getBotToken()
	if err != nil {
		panic(err.Error())
	}

	bot, err := newBot(botToken, BotUpdateTimeout)
	if err != nil {
		panic(err.Error())
	}

	browser, err := newBrowser()
	if err != nil {
		panic(err.Error())
	}

	updates, err := bot.GetUpdates()
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if !isUserAccessAllowed(update.Message.From.UserName) {
			userAppeal := "you"
			if update.Message.From.UserName != "" {
				userAppeal = fmt.Sprintf("@%s.", update.Message.From.UserName)
			}

			bot.ReplyToMessage(*update.Message, fmt.Sprintf("Access denied for %s.", userAppeal))
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

func getBotToken() (string, error) {
	if len(os.Args) >= 2 {
		return os.Args[1], nil
	}

	return "", errors.New("Bot token string must be provided as first argument.")
}

func isUserAccessAllowed(username string) bool {
	if len(os.Args) <= 2 {
		return true
	}

	for _, whitelistedUsername := range os.Args[2:] {
		if username == whitelistedUsername {
			return true
		}
	}
	return false
}
