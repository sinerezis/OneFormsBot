package main

import (
	bot "oneforms/OneFormsBot"
	sheets "oneforms/OneFormsSheets"

	//"oneforms/token"

	"fmt"
	"os"
	"strconv"

	"sync"
	"time"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

// Функция нужна для того,что бы неприрывно обрабатывать данные из таблицы
// Мы забираем копию выбранного лиcта таблицы
// Затем сканируем ее на предмет изменений
// И если они есть - направляем изменения через бота
// в чат

func SendOrders(bot *tgbotapi.BotAPI, sheetUrl string) {
	fmt.Println("work")
	for {
		sheet, _ := sheets.StartSheet(sheetUrl)
		orders, _ := sheets.CheckSheet(sheet)
		if len(orders) > 0 {
			for _, order := range orders {
				ChatId, _ := strconv.Atoi(os.Getenv("ChatId"))
				formatMessage := fmt.Sprintln("Новый заказ: ", order)
				fmt.Println(formatMessage, ChatId, "test")
				message := tgbotapi.NewMessage(int64(ChatId), formatMessage)
				bot.Send(message)
			}
		}
		time.Sleep(10 * time.Second)

	}
}

func main() {
	var wg sync.WaitGroup
	bot, _ := bot.NewBot(os.Getenv("Telegram_api_token"))

	wg.Add(1)
	go SendOrders(bot, os.Getenv("SheetURL"))

	wg.Wait()

}
