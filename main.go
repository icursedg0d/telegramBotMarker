package main

import (
	"app/tg/clients/telegram"
	"flag"
	"log"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	tgClient := telegram.New(tgBotHost, mustToken())
}

func mustToken() string {
	token := flag.String("tg-bot-token", "", "Telegram bot token")

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
