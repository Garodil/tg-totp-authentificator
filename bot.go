package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func LoginBot(botToken string) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true
	log.Printf("Bot авторизован как @%s", bot.Self.UserName)

	return bot
}

func HandleUpdates(updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI, secret string, chatId int64) {

	for {
		log.Println("Blocked cycle")
		update := <-updates
		log.Println("Received update")
		command := update.Message.Command()
		// if chatId != update.Message.Chat.ID {
		// 	continue
		// }
		switch command {
		case "start":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Используй /otp, чтобы получить свой код для входа в GitHub.")
			bot.Send(msg)
		case "otp":
			otpCode := generateTOTP(secret)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Твой TOTP код для GitHub: %s", otpCode))
			bot.Send(msg)
		}
	}
}
