package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var timeStep int = 30

func main() {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN is not set")
	}

	secret := os.Getenv("SECRET")
	if secret == "" {
		log.Fatal("SECRET is not set")
	}

	chat_id := os.Getenv("CHAT_ID")
	if chat_id == "" {
		log.Fatal("CHAT_ID is not set")
	}
	chatId, err := strconv.ParseInt(chat_id, 10, 64)
	if err != nil {
		log.Fatal("CHAT_ID is wrong")
	}

	webhookURL := os.Getenv("RENDER_EXTERNAL_URL") + "/webhook"
	if webhookURL == "/webhook" {
		log.Fatal("RENDER_EXTERNAL_URL is not set")
	}

	bot := LoginBot(botToken)

	var wg sync.WaitGroup
	wg.Add(1)

	go HandleUpdates(bot, webhookURL, secret, chatId, &wg)

	err = http.ListenAndServe("0.0.0.0:8443", nil)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
