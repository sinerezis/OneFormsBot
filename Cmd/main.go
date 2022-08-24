package main

import (
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

func SendOrders(sheetUrl string) error {
	fmt.Print(os.Getenv("Telegram_api_token"))
	bot, err := tgbotapi.NewBotAPI(os.Getenv("Telegram_api_token"))

	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Print(os.Getenv("Telegram_api_token"))
	if bot == nil {
		fmt.Println("bot is nil")

	}
	fmt.Println("work")
	for {
		sheet, _ := sheets.StartSheet(sheetUrl)
		orders, _ := sheets.CheckSheet(sheet)
		if len(orders) > 0 {
			for _, order := range orders {
				ChatId, _ := strconv.Atoi(os.Getenv("ChatId"))
				formatMessage := fmt.Sprintln("Новый заказ: ", order)
				fmt.Println(formatMessage, ChatId, "test")
				fmt.Println(os.Getenv("SheetURL"), "testurl")
				message := tgbotapi.NewMessage(int64(ChatId), formatMessage)
				fmt.Println("message created")

				bot.Send(message)
				fmt.Println("message sended")
			}
		}
		time.Sleep(10 * time.Second)

	}
}

func main() {
	var wg sync.WaitGroup
	//bot, _ := bot.NewBot(os.Getenv("Telegram_api_token"))

	wg.Add(1)
	go SendOrders(os.Getenv("SheetURL"))

	wg.Wait()

}
