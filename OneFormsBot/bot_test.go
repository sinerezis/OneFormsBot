package oneformsbot

import (
	testtoken "oneforms/test_token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBot(t *testing.T) {

	bot, error := NewBot(testtoken.Telegram_api_token)
	assert.NoError(t, error)
	assert.NotNil(t, bot)
}
