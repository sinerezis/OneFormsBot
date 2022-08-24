package main

import (
	"log"
	"net/http"
	bot "oneforms/OneFormsBot"
	sheets "oneforms/OneFormsSheets"
	"oneforms/token"

	"fmt"
	"os"

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
	log.Print("Начинаем сканирование!")
	for {
		sheet, _ := sheets.StartSheet(sheetUrl)
		orders, _ := sheets.CheckSheet(sheet)

		if len(orders) > 0 {
			for _, order := range orders {

				formatMessage := fmt.Sprintln("Новый заказ: ", order)
				message := tgbotapi.NewMessage(int64(token.ChatId), formatMessage)

				bot.Send(message)
				log.Printf("Заказ %s отправлен в чат", order)
			}
		}
		time.Sleep(10 * time.Second)

	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	http.HandleFunc("/", bot.MainHandler)
	go http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	go SendOrders(os.Getenv("SheetURL"))

	wg.Wait()

}
