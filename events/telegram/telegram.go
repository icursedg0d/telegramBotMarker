package telegram

import "app/tg/clients/telegram"

type Processor struct {
	tg     *telegram.Client
	offset int
	// storage
}
