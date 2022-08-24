package main

import (
	"net/http"
	bot "oneforms/OneFormsBot"
	sheets "oneforms/OneFormsSheets"

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
	bot, err := bot.NewBot()
	if err != nil {
		return err
	}
	fmt.Println("work")
	for {
		sheet, _ := sheets.StartSheet(sheetUrl)
		orders, _ := sheets.CheckSheet(sheet)

		if len(orders) > 0 {
			for _, order := range orders {

				ChatId, _ := strconv.Atoi(os.Getenv("ChatId"))
				formatMessage := fmt.Sprintln("Новый заказ: ", order)
				message := tgbotapi.NewMessage(int64(ChatId), formatMessage)

				bot.Send(message)
			}
		}
		time.Sleep(10 * time.Second)

	}
}

func main() {
	fmt.Println("1")
	var wg sync.WaitGroup

	wg.Add(2)

	http.HandleFunc("/", bot.MainHandler)
	go http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	go SendOrders(os.Getenv("SheetURL"))

	wg.Wait()

}
