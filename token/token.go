package token

import (
	"os"
	"strconv"
	"strings"
)

var Telegram_api_token = strings.Trim(os.Getenv("Telegram_api_token"), "\n")
var SheetURL string = os.Getenv("SheetURL")
var ChatId, _ = strconv.Atoi(os.Getenv("ChatId"))
