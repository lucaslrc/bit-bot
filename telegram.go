package main

import (
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func initTelegramService(telegramToken string) {

	// apiNews := JSONConfig.Config.NewsAPIToken
	apiWeather := JSONConfig.Config.OpenWeatherMAPToken
	city := JSONConfig.Config.CityName
	state := JSONConfig.Config.StateCode
	country := JSONConfig.Config.CountryCode

	bot, err := tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		log.Panic(err)
		return
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	for {
		time.Sleep(5 * time.Second)

		group := tgbotapi.NewMessage(-1001421497962, getWeather(apiWeather, city, state, country))

		bot.Send(group)
	}

	// u := tgbotapi.NewUpdate(0)
	// u.Timeout = 5

	// updates, err := bot.GetUpdatesChan(u)

	// for update := range updates {
	// 	if update.Message == nil { // ignore any non-Message Updates
	// 		continue
	// 	}

	// 	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	// 	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	// 	msg.ReplyToMessageID = update.Message.MessageID

	// 	bot.Send(msg)
	// }
}
