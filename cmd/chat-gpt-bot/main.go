package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"modules/internal/tgbot"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(tgbot.BotToken)
	if err != nil {
		log.Fatalf("Error creating Telegram tgbot: %v", err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		tgbot.HandleTelegramUpdate(update, bot)
	}
}
