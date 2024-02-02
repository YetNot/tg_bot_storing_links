package telegram

import "tg_bot_storing_links/clients/telegram"

type Processor struct {
	tg     *telegram.Client
	offset int
}

func New(client *telegram.Client) {
}
