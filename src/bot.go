package main

import (
	"github.com/Syfaro/telegram-bot-api"
)

const BotUpdateTimeout = 60

type Bot struct {
	api    tgbotapi.BotAPI
	config tgbotapi.UpdateConfig
}

func newBot(token string, timeout int) (*Bot, error) {
	config := tgbotapi.NewUpdate(0)
	config.Timeout = timeout

	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &Bot{*api, config}, nil
}

func (bot *Bot) GetUpdates() (tgbotapi.UpdatesChannel, error) {
	return bot.api.GetUpdatesChan(bot.config)
}

func (bot *Bot) sendMessage(chat tgbotapi.Chat, msgText string) error {
	message := tgbotapi.NewMessage(chat.ID, msgText)

	_, err := bot.api.Send(message)
	return err
}

func (bot *Bot) ReplyToMessage(message tgbotapi.Message, replyText string) error {
	replyMessage := tgbotapi.NewMessage(message.Chat.ID, replyText)
	replyMessage.ReplyToMessageID = message.MessageID

	_, err := bot.api.Send(replyMessage)
	return err
}
