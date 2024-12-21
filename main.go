package main

import (
	tgClient "app/tg/clients/telegram"
	event_consumer "app/tg/consumer/event-consumer"
	"app/tg/events/telegram"
	"app/tg/storage/files"
	"flag"
	"log"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files-storage"
	batchSize   = 100
)

func main() {

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath))
	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String("tg-bot-token", "", "Telegram bot token")
	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
