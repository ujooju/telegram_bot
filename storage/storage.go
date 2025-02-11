package storage

import (
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type Database struct {
	Url    string
	tunnel *sqlx.DB
}

func NewDatabase() *Database {
	databaseURL := os.Getenv("STORAGE_URL")
	if databaseURL == "" {
		log.Error().Msg("check if STORAGE_URL is set")
	}
	return &Database{
		Url: databaseURL,
	}
}

func (db *Database) StartDatabase() error {
	var err error
	db.tunnel, err = sqlx.Connect("pgx", db.Url)
	if err != nil {
		return err
	}
	db.InitSchema()
	return nil
}
