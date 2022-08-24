package oneformsbot

import (
	"log"
	"os"
	"strings"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

// Инициализируем и запускаем бота
func NewBot() (*tgbotapi.BotAPI, error) {
	key := os.Getenv("Telegram_api_token")
	key = strings.Trim(key, "\n")
	bot, err := tgbotapi.NewBotAPI(key)

	if err != nil {
		return nil, err
	}
	log.Printf("Connect to %s", bot.Self.UserName)
	return bot, nil

}
