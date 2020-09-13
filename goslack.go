package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ilho-tiger/goslack/slack"
)

func main() {
	messageFlagPtr := flag.String("message", "", "a message to send to Slack Incoming Webhook")
	webhookAddrFlagPtr := flag.String("url", "", "(optional) a Slack Incoming Webhook URL to use. If not present the environment variable `slack_webhook` will be used")
	flag.Parse()

	if messageFlagPtr == nil || *messageFlagPtr == "" {
		flag.Usage()
		fmt.Println("")
		log.Fatalln("error: no message data provided")
	}

	if webhookAddrFlagPtr != nil && *messageFlagPtr != "" {
		slack.SetWebhookURL(*webhookAddrFlagPtr)
	}
	slack.SendMessage(*messageFlagPtr)
}
