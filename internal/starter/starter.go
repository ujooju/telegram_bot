package starter

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/ujooju/telegram_bot/internal/webhookServer"
	"github.com/ujooju/telegram_bot/storage"
)

type Bot struct {
	token   string
	db      *storage.Database
	webhook *webhookServer.Webhook
}

func Start() {
	bot := newBot()
	log.Info().Msg("bot started listening " + bot.webhook.Url)
}

func newBot() *Bot {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Error().Msg("check if TELEGRAM_BOT_TOKEN is set")
	}
	db := storage.NewDatabase()
	err := db.StartDatabase()
	if err != nil {
		log.Error().Err(err)
	}
	webhook := webhookServer.NewWebhook(token)
	bot := &Bot{
		token:   token,
		db:      db,
		webhook: webhook,
	}
	bot.webhook.Start()
	return bot
}
