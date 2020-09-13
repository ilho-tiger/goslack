package slack

import (
	"testing"
)

func Test_getWebhookURL(t *testing.T) {
	err := setWebhookURLFromEnvIfEmptyy()
	t.Log("value:", webhookURL)
	t.Log("err:", err)
}
