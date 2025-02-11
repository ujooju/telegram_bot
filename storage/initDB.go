package storage

import (
	"io"
	"os"

	"github.com/rs/zerolog/log"
)

func (db *Database) InitSchema() {
	fileConfigRequest, err := os.Open("config/storage/schema.txt")
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	defer fileConfigRequest.Close()
	schemaRequestByte, err := io.ReadAll(fileConfigRequest)
	if err != nil {
		log.Error().Msg(err.Error())
	}
	shemaRequest := string(schemaRequestByte)
	_, err = db.tunnel.Query(shemaRequest)
	if err != nil {
		log.Error().Msg(err.Error())
	}
}
