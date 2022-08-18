package oneformsbot

import (
	"log"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

// Инициализируем и запускаем бота
func NewBot(token string) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Connect to %s", bot.Self.UserName)
	return bot

}
