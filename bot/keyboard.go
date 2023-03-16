package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var getDataKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Мой первый запрос"),
		tgbotapi.NewKeyboardButton("Все запросы"),
	),
)

// var key = tgbotapi.NewKeyboardButton("Мой первый запрос")

// var getDataKeyboard = tgbotapi.NewReplyKeyboard(
//     tgbotapi.NewKeyboardButtonRow(
//         tgbotapi.NewKeyboardButton("Мой первый запрос"),
//         tgbotapi.NewKeyboardButton("Все запросы"),
//     ),
// )

// var getDataKeyboard = tgbotapi.NewReplyKeyboard(
// 	tgbotapi.NewKeyboardButtonRow(
// 		tgbotapi.NewKeyboardButton("Мой первый запрос"),
// 		tgbotapi.NewKeyboardButton("Все запросы"),
// 	),
// )
