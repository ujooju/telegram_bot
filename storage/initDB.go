package storage

import (
	"io"
	"os"

	"github.com/rs/zerolog/log"
)

func initDB() {
	fileConfigRequest, err := os.Open("config/request.txt")
	if err != nil {
		log.Error().Err(err)
	}
	configRequest, err := io.ReadAll(fileConfigRequest)
	if err != nil {
		log.Fatal().Err(err)
	}

}
