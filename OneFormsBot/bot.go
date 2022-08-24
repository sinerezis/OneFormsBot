package oneformsbot

import (
	"log"
	"net/http"
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

func MainHandler(resp http.ResponseWriter, _ *http.Request) {
	resp.Write([]byte("Привет, это бот One Froms"))
}
