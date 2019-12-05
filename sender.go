package discordhook

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type Sender struct {
	WebhookURL string
}

func NewSender(webhookUrl string) Sender {
	return Sender{WebhookURL: webhookUrl}
}

func (s *Sender) Send(message DiscordMessage) error {
	client := http.Client{
		Timeout: time.Second * 15,
	}

	requestJson, err := json.Marshal(message)

	request, err := http.NewRequest("POST", s.WebhookURL, bytes.NewBuffer(requestJson))
	request.Header.Set("Content-type", "application/json")

	response, err := client.Do(request)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	if err != nil {
		return err
	}

	return nil
}