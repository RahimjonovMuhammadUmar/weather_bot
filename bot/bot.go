package bot

import (
	"log"

	"github.com/RahimjonovMuhammadUmar/weather_bot/config"
	"github.com/RahimjonovMuhammadUmar/weather_bot/storage"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotHandler struct {
	cfg     config.Config
	storage storage.StorageI
	bot     *tgbotapi.BotAPI
}

func New(cfg config.Config, strg storage.StorageI) BotHandler {
	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	return BotHandler{
		cfg:     cfg,
		storage: strg,
		bot:     bot,
	}
}

func (h *BotHandler) Start() {
	log.Printf("Authorized on account %s", h.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := h.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			go h.HandleBot(update)
		}
	}

}

func (h *BotHandler) HandleBot(update tgbotapi.Update) {
	user := &storage.User{
		TgID:      update.Message.From.ID,
		Firstname: update.Message.From.FirstName,
		Lastname:  update.Message.From.LastName,
	}
	var err error
	if update.Message.Command() == "start" {
		exist := h.CheckUserExistence(user)
		if exist {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "С возвращением!")
			msg.ReplyMarkup = getDataKeyboard
			if _, err = h.bot.Send(msg); err != nil {
				log.Println(err)
			}
		} else {
			err = h.WelcomeMessage(user)
		}
	} else if update.Message.Text != "" {
		switch update.Message.Text {
		case "Мой первый запрос":
			err = h.GetFirstRequest(update, user)
		case "Все запросы":
			err = h.GetAllRequest(update, user)
		default:
			err = h.GetTemperature(update, user)
		}
	}
	if err != nil {
		log.Println("failed to handle message: ", err)
		h.SendMessage(user, "Произошла ошибка, пожалуйста перезапустите бота /start")
	}
}
