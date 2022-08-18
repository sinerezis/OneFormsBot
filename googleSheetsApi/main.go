package main

import (
	bot "oneforms/cmd/OneFormsBot"
	"oneforms/token"

	"context"
	"fmt"
	"io/ioutil"
	"log"

	"strconv"
	"sync"
	"time"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"
)

// Инициализируем доступ к таблице
func startSheet() (*spreadsheet.Sheet, error) {
	data, err := ioutil.ReadFile(token.Path_To_Client_Secret)
	if err != nil {
		log.Fatal(err)
	}

	conf, err := google.JWTConfigFromJSON(data, spreadsheet.Scope)
	if err != nil {
		return nil, err
	}

	client := conf.Client(context.TODO())
	service := spreadsheet.NewServiceWithClient(client)
	spreadsheet, err := service.FetchSpreadsheet(token.SheetURL)
	if err != nil {
		return nil, err
	}

	sheet, err := spreadsheet.SheetByID(0)
	if err != nil {
		return nil, err
	}

	return sheet, nil

}

// Эта функция проверяет, отличается ли текущее кол-во строк в таблице от
// записанного. Если отличается - мы проходимся по каждой строке и записываем ее в массив.
// увеличивая счетчик, записанный в таблице при каждой итерации.
// Как только счетчики сравниваются - мы выходим из программы
func ChekSheet(sheet *spreadsheet.Sheet) ([]string, error) {
	//Инициализируем счетчик - текущая длина таблицы
	counter := len(sheet.Rows)

	// Иницилизируем срез для хранения новых заказов
	var orders []string

	//Бесконечный цикл
	for {

		// Сканируем число прочитанных элементов, которое записано в самой таблице
		countOfRows, _ := strconv.Atoi(sheet.Rows[0][20].Value)

		// Если все заказы из таблицы прочитаны - выходим
		if countOfRows == counter {
			return orders, nil
		}

		// Если в таблице есть новые заказы
		// и если значение в графе "рамер" не пустое -
		// добавляем заказ  в срез
		// и увеличиваем счетчик прочитанных заказов на 1
		if counter > countOfRows {
			for ; countOfRows <= counter; countOfRows++ {
				if sheet.Rows[countOfRows][3].Value != "" {
					orders = append(orders, sheet.Rows[countOfRows][3].Value)

				}

				// Обновляем кол-во прочитанных заказов, записаное
				// в таблице
				sheet.Update(0, 20, strconv.Itoa(countOfRows))
				err := sheet.Synchronize()
				if err != nil {
					return orders, nil
				}
			}
		}
		// Если из таблицы удаляют значение - переписываем значение счетчика
		// прочитанных значений
		if countOfRows > counter {
			sheet.Update(0, 20, strconv.Itoa(counter))
			err := sheet.Synchronize()
			if err != nil {
				return orders, nil
			}
		}
	}
}

// Функция нужна для того,что бы неприрывно обрабатывать данные из таблицы
func SendOrders(bot *tgbotapi.BotAPI) {
	fmt.Println("work")
	for {
		sheet, _ := startSheet()
		orders, _ := ChekSheet(sheet)
		if len(orders) > 0 {
			for _, order := range orders {
				message := tgbotapi.NewMessage(token.ChatId, order)
				bot.Send(message)
			}
		}
		time.Sleep(10 * time.Second)

	}
}

func main() {
	var wg sync.WaitGroup
	bot := bot.NewBot(token.Telegram_api_token)

	wg.Add(1)
	go SendOrders(bot)

	wg.Wait()

}
