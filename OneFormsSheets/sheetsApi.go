package oneformssheets

import (
	"fmt"
	"io/ioutil"
	"os"

	"context"

	"strconv"

	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"
)

// Инициализируем доступ к таблице
func StartSheet(sheetUrl string) (*spreadsheet.Sheet, error) {

	data, err := ioutil.ReadFile("tmp/build_1426e25c/google-credentials.json")
	if err != nil {
		return nil, err
	}

	// Генерируем конфиг из прочитанного
	// json файла
	conf, err := google.JWTConfigFromJSON(data, spreadsheet.Scope)
	if err != nil {
		return nil, fmt.Errorf("1")
	}

	// Создаем клиента
	client := conf.Client(context.TODO())
	// Создаем сервис
	service := spreadsheet.NewServiceWithClient(client)
	// Присоединяемся к таблице по ее токену
	spreadsheet, err := service.FetchSpreadsheet(os.Getenv("SheetURL"))
	if err != nil {
		return nil, err
	}

	// Читаем первый лист таблицы
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
func CheckSheet(sheet *spreadsheet.Sheet) ([]string, error) {
	//Инициализируем счетчик - текущая длина таблицы
	counter := len(sheet.Rows)

	// Иницилизируем срез для хранения новых заказов
	var orders []string

	//Бесконечный цикл
	for {

		// Сканируем число прочитанных элементов, которое записано в самой таблице
		countOfRows, err := strconv.Atoi(sheet.Rows[0][20].Value)
		if err != nil {
			return orders, err
		}

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
					return orders, err
				}

			}
		}
		// Если из таблицы удаляют значение - переписываем значение счетчика
		// прочитанных значений
		if countOfRows > counter {
			sheet.Update(0, 20, strconv.Itoa(counter))
			err := sheet.Synchronize()
			if err != nil {
				return orders, err
			}
		}
	}
}
