package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var webhookURL string = ""

const (
	webhookEnvVarName = "slack_webhook"
)

// SendMessage will send given message as a Slack incoming webhook
func SendMessage(message string) error {
	setWebhookURLFromEnvIfEmptyy()
	var err error = nil

	requestBody, err := json.Marshal(map[string]string{
		"text": message,
	})

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("Failed to send message: %s", resp.Status)
	}
	defer resp.Body.Close()
	return err
}

// SetWebhookURL sets the webhook url to use for Slack
func SetWebhookURL(url string) {
	webhookURL = url
}

// WebhookURL returns the currently configured webhookURL value
func WebhookURL() string {
	setWebhookURLFromEnvIfEmptyy()
	return webhookURL
}

func setWebhookURLFromEnvIfEmptyy() error {
	if webhookURL != "" {
		return nil // url has been already set with value
	}
	return setWebhookURLFromEnv()
}

func setWebhookURLFromEnv() error {
	var err error
	var value string

	value = os.Getenv(webhookEnvVarName)
	if value != "" {
		webhookURL = value
	} else {
		err = fmt.Errorf("No environment variable named '%s' found", webhookEnvVarName)
	}
	return err
}
