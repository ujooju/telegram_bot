package starter

import (
	"log"
	"os"
)

type Bot struct {
	token string
}

func Start() {
	bot := newBot()
}

func newBot() *Bot {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("check that TELEGRAM_BOT_TOKEN is set")
	}
	bot := &Bot{
		token: token,
	}
	return bot
}
