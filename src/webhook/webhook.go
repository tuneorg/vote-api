package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Embed struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	URL         string `json:"url,omitempty"`
	Color       int    `json:"color,omitempty"`
}

type WebhookPayload struct {
	Embeds []Embed `json:"embeds"`
}

func SendWebhook(webhookURL string, embed Embed) error {
	payload := WebhookPayload{
		Embeds: []Embed{embed},
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Accept 200 and 204 as successful responses
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to send webhook, status code: %d", resp.StatusCode)
	}

	return nil
}
