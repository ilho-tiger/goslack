package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var webhookURL string = ""

const (
	enableSlackEnvVarName = "enable_slack"
	webhookEnvVarName     = "slack_webhook"
)

// SendMessage will send given message as a Slack incoming webhook
func SendMessage(message string) error {
	if IsSlackEnabled() == false {
		log.Println("Slack message is disabled by the envrionment, skipping.")
		fmt.Printf("\n%s\n", message)
		return nil
	}
	getWebhookURL()
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

// IsSlackEnabled returns slack message configuration
func IsSlackEnabled() bool {
	value := os.Getenv(enableSlackEnvVarName)
	if value == "true" {
		return true
	}
	return false
}

// SetWebhookURL sets the webhook url to use for Slack
func SetWebhookURL(url string) {
	webhookURL = url
}

func getWebhookURL() error {
	if webhookURL != "" {
		return nil // url has been already set with value
	}
	return getWebhookURLFromEnv()
}

func getWebhookURLFromEnv() error {
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
