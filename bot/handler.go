package bot

import (
	"fmt"
	"log"
	"strings"

	"github.com/RahimjonovMuhammadUmar/weather_bot/storage"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/cast"
)

func (h *BotHandler) CheckUserExistence(user *storage.User) bool {
	exists, err := h.storage.CheckUserExistence(user.TgID)
	if err != nil {
		log.Printf("CheckUserExistence failed to check info in db for user %d: %s\n", user.TgID, err)
	}

	return exists
}

func (h *BotHandler) WelcomeMessage(user *storage.User) error {
	msg := tgbotapi.NewMessage(user.TgID, "Добро пожаловать!\nВведите имя города: ")
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	if _, err := h.bot.Send(msg); err != nil {
		log.Printf("Failed to send welcoming text to user %d: %s\n", user.TgID, err)
	}
	_, err := h.storage.CreateUser(user)
	if err != nil {
		log.Printf("CreateUser failed to create user in db for user %d: %s\n", user.TgID, err)
	}
	return err
}

func (h *BotHandler) GetTemperature(update tgbotapi.Update, user *storage.User) error {
	city := update.Message.Text
	temperature := getCityTemperature(city, h.cfg)
	if temperature != nil {
		responseText := fmt.Sprintf("Температура в %s %.1f°C", city, *temperature)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, responseText)
		msg.ReplyMarkup = getDataKeyboard
		_, err := h.bot.Send(msg)
		if err != nil {
			log.Printf("Failed to send temperature to user %d: %s\n", user.TgID, err)
		}
		user.City = city
		if _, err := h.storage.CreateRequest(user); err != nil {
			log.Printf("CreateRequest failed to create request in db for request %s: %s\n", user.City, err)
		}
	} else {
		responseText := fmt.Sprintf("Извините, нам не удалось найти температуру для %s. Пожалуйста, проверьте правописание и повторите попытку.", update.Message.Text)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, responseText)
		_, err := h.bot.Send(msg)
		if err != nil {
			log.Printf("Failed to send warning to user %d: %s\n", user.TgID, err)
		}
	}
	return nil
}

func (h *BotHandler) GetFirstRequest(update tgbotapi.Update, user *storage.User) error {
	firstRequest, err := h.storage.GetFirstRequest(user.TgID)
	if err != nil {
		log.Printf("GetFirstRequest failed to get info from db for user %d: %s\n", user.TgID, err)
	}
	response := fmt.Sprintf(" - %s\n   %s\n\n", firstRequest.City, firstRequest.CreatedAt)

	msg := tgbotapi.NewMessage(user.TgID, response)
	if _, err := h.bot.Send(msg); err != nil {
		log.Printf("GetFirstRequest failed to send result for user %d: %s\n", user.TgID, err)
	}

	return nil
}

func (h *BotHandler) GetAllRequest(update tgbotapi.Update, user *storage.User) error {
	allRequests, err := h.storage.GetAllRequests(user.TgID)
	if err != nil {
		log.Printf("GetAllRequests failed to get info from db for user %d: %s\n", user.TgID, err)
	}
	var builder strings.Builder
	for _, req := range allRequests {
		builder.WriteString(fmt.Sprintf(" - %s\n   %s\n\n", req.City, req.CreatedAt))
	}
	msg := tgbotapi.NewMessage(user.TgID, cast.ToString(builder.String()))
	if _, err := h.bot.Send(msg); err != nil {
		log.Printf("Failed to send message to user %d: %s\n", user.TgID, err)
	}

	return nil
}

func (h *BotHandler) SendMessage(user *storage.User, message string) {
	msg := tgbotapi.NewMessage(user.TgID, message)
	if _, err := h.bot.Send(msg); err != nil {
		log.Printf("Failed to send message to user %d: %s\n", user.TgID, err)
	}
}
