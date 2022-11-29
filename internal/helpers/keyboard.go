package helpers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var StartKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/help"),
		tgbotapi.NewKeyboardButton("/verbs"),
	),
)

var VerbsInitKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/study"),
		tgbotapi.NewKeyboardButton("/test_letter"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/test"),
		tgbotapi.NewKeyboardButton("/back"),
	),
)

var VerbsTestKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/reroll"),
		tgbotapi.NewKeyboardButton("/answers"),
		tgbotapi.NewKeyboardButton("/back"),
	),
)

var StudyKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("а - б"),
		tgbotapi.NewKeyboardButton("г - з"),
		tgbotapi.NewKeyboardButton("и - ө"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("с - т"),
		tgbotapi.NewKeyboardButton("у - х"),
		tgbotapi.NewKeyboardButton("ц - я"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/back"),
	),
)
