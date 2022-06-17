package main

import (
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/kelseyhightower/envconfig"
)

const (
	PushoverURL = "https://api.pushover.net/1/messages.json"
)

type Config struct {
	PushoverAPIToken string `envconfig:"PUSHOVER_API_TOKEN"`
	PushoverUserKey  string `envconfig:"PUSHOVER_USER_KEY"`
}

func main() {
	var c Config
	err := envconfig.Process("", &c)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.PostForm(PushoverURL, url.Values{
		"token":   {c.PushoverAPIToken},
		"user":    {c.PushoverUserKey},
		"message": {"Your task is done!"},
	})

	if err != nil {
		log.Fatalf("Error sending message: %s", err)
	}
	defer resp.Body.Close()

	b, _ := io.ReadAll(resp.Body)
	log.Printf("Response: %s", b)
}
