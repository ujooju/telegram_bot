package webhookserver

import (
	"log"
	"os"
)

type Webhook struct {
	url string
}

func (webhook *Webhook) Start() {

}

type WebhookParams struct {
	url string
}

func startWebhook() {

}

func newWebhook() *Webhook {
	url := os.Getenv("WEBHOOK_URL")
	if url == "" {
		log.Fatal("check that WEBHOOK_URL is set")
	}
	webhook := &Webhook{
		url: url,
	}
	return webhook
}
