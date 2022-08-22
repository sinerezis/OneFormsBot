package main

import (
	bot "oneforms/OneFormsBot"
	sheets "oneforms/OneFormsSheets"
	"oneforms/token"

	"fmt"

	"sync"
	"time"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

// Функция нужна для того,что бы неприрывно обрабатывать данные из таблицы
// Мы забираем копию выбранного лиcта таблицы
// Затем сканируем ее на предмет изменений
// И если они есть - направляем изменения через бота
// в чат

func SendOrders(bot *tgbotapi.BotAPI, pathToToken string, sheetUrl string) {
	fmt.Println("work")
	for {
		sheet, _ := sheets.StartSheet(pathToToken, sheetUrl)
		orders, _ := sheets.CheckSheet(sheet)
		if len(orders) > 0 {
			for _, order := range orders {
				formatMessage := fmt.Sprintln("Новый заказ: ", order)
				message := tgbotapi.NewMessage(token.ChatId, formatMessage)
				bot.Send(message)
			}
		}
		time.Sleep(10 * time.Second)

	}
}

func main() {
	var wg sync.WaitGroup
	bot, _ := bot.NewBot(token.Telegram_api_token)

	wg.Add(1)
	go SendOrders(bot, token.Path_To_Client_Secret, token.SheetURL)

	wg.Wait()

}
