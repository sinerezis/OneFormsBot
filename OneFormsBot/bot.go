package oneformsbot

import (
	config "oneforms/config"

	"log"
	"net/http"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

// Инициализируем и запускаем бота
func NewBot() (*tgbotapi.BotAPI, error) {

	bot, err := tgbotapi.NewBotAPI(config.Telegram_api_token)

	if err != nil {
		return nil, err
	}

	go func() {
		updates := bot.ListenForWebhook("/" + bot.Token)
		for update := range updates {
			log.Print(update.Message.From, update.Message)
		}
	}()

	log.Printf("Connect to %s", bot.Self.UserName)
	return bot, nil

}

func MainHandler(resp http.ResponseWriter, _ *http.Request) {
	resp.Write([]byte("Привет, это бот One Forms!"))
}
