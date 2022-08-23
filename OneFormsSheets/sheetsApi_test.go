package oneformssheets

import (
	//testtoken "oneforms/test_token"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartSheet(t *testing.T) {
	sheet, err := StartSheet(testtoken.Google_Sheets_api_token, testtoken.SheetURL)
	assert.NoError(t, err)
	assert.NotNil(t, sheet)

}

func TestCheckLenSheet(t *testing.T) {
	sheet, err := StartSheet(testtoken.Google_Sheets_api_token, testtoken.SheetURL)
	assert.NoError(t, err)

	lenOfSheet := len(sheet.Rows)
	recordedValue, _ := strconv.Atoi(sheet.Rows[0][20].Value)

	// Если в таблице все данные прочитаны - она
	// должна вернуть пустой срез
	if lenOfSheet == recordedValue {

		returnedValue, _ := CheckSheet(sheet)
		assert.Equal(t, len(returnedValue), 0)
	}

	// Если в таблице есть новые данные - она
	// должна вернуть срез длинной в кол-во
	// новых строк в таблице
	if lenOfSheet > recordedValue {
		returnedValue, _ := CheckSheet(sheet)
		assert.Equal(t, len(returnedValue), (lenOfSheet - recordedValue))
		recordedValue, _ = strconv.Atoi(sheet.Rows[0][20].Value)
		assert.Equal(t, recordedValue, lenOfSheet)
	}

	// Если из таблицы были удалены заказы
	// Она должна вернуть пустой срез
	// и перезаписать значение прочитанных заказов
	if lenOfSheet < recordedValue {
		returnedValue, _ := CheckSheet(sheet)
		assert.Equal(t, len(returnedValue), 0)
		recordedValue, _ = strconv.Atoi(sheet.Rows[0][20].Value)
		assert.Equal(t, recordedValue, lenOfSheet)

	}

}
