package main

import (
	"fmt"
	"log"
	"sync"

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

func HandleUpdates(bot *tgbotapi.BotAPI, webhookURL string, secret string, chatId int64, wg *sync.WaitGroup) {
	updates := bot.ListenForWebhook(webhookURL)
	wg.Done()
	log.Println("Listening for updates on " + webhookURL)
	webhookInfo, err := bot.GetWebhookInfo()
	if err != nil {
		log.Println("Error receiveng webhook info")
		return
	}
	log.Println("Listening for updates on (2)" + webhookInfo.URL)
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
