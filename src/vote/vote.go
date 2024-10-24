package vote

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	config "vote-api/src"
	"vote-api/src/webhook"
)

type VotePayload struct {
	Bot       string `json:"bot"`
	User      string `json:"user"`
	Type      string `json:"type"`
	IsWeekend bool   `json:"isWeekend"`
}

func CheckAuthorization(r *http.Request, conf config.Config) bool {
	expectedAuthKey := conf.TOPGG_KEY
	receivedAuthKey := r.Header.Get("Authorization")
	return receivedAuthKey == expectedAuthKey
}

func VoteHandler(conf config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !CheckAuthorization(r, conf) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			log.Println("Unauthorized request received.")
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Can't read body", http.StatusBadRequest)
			return
		}

		var payload VotePayload
		err = json.Unmarshal(body, &payload)
		if err != nil {
			log.Printf("Error unmarshalling JSON: %v", err)
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		log.Printf("Received vote from user %s for bot %s, isWeekend: %v", payload.User, payload.Bot, payload.IsWeekend)

		embed := webhook.Embed{
			Title:       conf.EMBED.Title,
			Description: replacePlaceholders(conf.EMBED.Description, payload),
			Color:       conf.EMBED.Color,
		}

		err = webhook.SendWebhook(conf.WEBHOOK_URL, embed)
		if err != nil {
			log.Printf("Error sending webhook: %v", err)
			http.Error(w, "Failed to send webhook", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Vote received"))
	}
}

func replacePlaceholders(description string, payload VotePayload) string {
	replacements := map[string]string{
		"{{user}}":    payload.User,
		"{{bot}}":     payload.Bot,
		"{{weekend}}": fmt.Sprintf("%t", payload.IsWeekend),
	}

	for placeholder, value := range replacements {
		description = strings.ReplaceAll(description, placeholder, value)
	}

	return description
}
