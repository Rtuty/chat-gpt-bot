package tgbot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"modules/internal/gpt"
	"os"
)

var BotToken = os.Getenv("TG_BOT_TOKEN")

func HandleTelegramUpdate(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.Message == nil || update.Message.Text == "" {
		return
	}

	userInput := update.Message.Text

	// Запрос к API ChatGPT
	chatGPTResponse, err := gpt.GetChatGPTResponse(userInput)
	if err != nil {
		log.Printf("Error getting ChatGPT response: %v", err)
		return
	}

	// Отправляем ответ пользователю
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, chatGPTResponse)
	_, err = bot.Send(msg)
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
}
