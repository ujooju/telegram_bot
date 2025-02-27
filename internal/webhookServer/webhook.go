package webhookServer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/rs/zerolog/log"
)

type Webhook struct {
	Url   string
	Token string
}

type WebhookParams struct {
	Url         string `json:"url"`
	Certificate string `json:"certificate"`
}

func (webhook *Webhook) Start() {

	keyFile, err := os.Open("key.pem")
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	keyInBytes, err := io.ReadAll(keyFile)
	keyFile.Close()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	keyContent := string(keyInBytes)

	params := WebhookParams{
		Url:         webhook.Url,
		Certificate: keyContent,
	}
	paramsEncoded, err := json.Marshal(params)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	fmt.Println(webhook.Token)
	resp, err := http.Post("https://api.telegram.org/bot"+webhook.Token+"/setWebhook", "application/json", bytes.NewReader(paramsEncoded))
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	defer resp.Body.Close()
	respContent, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	fmt.Println(string(respContent))

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	err = http.ListenAndServeTLS(":443", "cert.pem", "key.pem", mux)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	log.Info().Msg("webhook started")
}

func NewWebhook(token string) *Webhook {
	url := os.Getenv("WEBHOOK_URL")
	if url == "" {
		log.Fatal().Msg("check if WEBHOOK_URL is set")
	}
	webhook := &Webhook{
		Url:   url,
		Token: token,
	}
	return webhook
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("incoming request")
	requestContent, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	log.Debug().Msg(string(requestContent))
	w.Write([]byte("hi"))
}
