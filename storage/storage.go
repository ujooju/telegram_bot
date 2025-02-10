package storage

import (
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type Database struct {
	url    string
	tunnel *sqlx.DB
}

func newDatabase() *Database {
	databaseURL := os.Getenv("STORAGE_URL")
	if databaseURL == "" {
		log.Fatal().Msg("chech if STORAGE_URL is set")
	}
	return &Database{
		url: databaseURL,
	}
}

func (db *Database) startDatabase() error {
	var err error
	db.tunnel, err = sqlx.Connect("pgx", db.url)
	if err != nil {
		return err
	}
	return nil
}
