package main

import (
	"app/tg/clients/telegram"
	"flag"
	"log"
)

// https://www.youtube.com/watch?v=4-jDamyg7QM&list=PLFAQFisfyqlWDwouVTUztKX2wUjYQ4T3l&index=4

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
