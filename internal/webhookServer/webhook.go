package webhookserver

import (
	"log"
	"os"
)

type WebHook struct {
	url string
}

func startWebHook {
	
}

func newWebHook() *WebHook {
	url := os.Getenv("WEBHOOK_URL")
	if url == "" {
		log.Fatal("check that WEBHOOK_URL is set")
	}
	webhook := &WebHook{
		url: url,
	}
	return webhook
}