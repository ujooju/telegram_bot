package starter

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/ujooju/telegram_bot/storage"
)

type Bot struct {
	token string
	db    *storage.Database
}

func Start() {
	bot := newBot()
	log.Info().Msg("connected to database at " + bot.db.Url)
}

func newBot() *Bot {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Error().Msg("check that TELEGRAM_BOT_TOKEN is set")
	}
	db := storage.NewDatabase()
	err := db.StartDatabase()
	if err != nil {
		log.Error().Err(err)
	}
	bot := &Bot{
		token: token,
		db:    db,
	}
	return bot
}
