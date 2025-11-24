package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/thanpawatpiti/notify"
	"github.com/thanpawatpiti/notify/providers/discord"
	"github.com/thanpawatpiti/notify/providers/line"
	"github.com/thanpawatpiti/notify/providers/msteams"
	"github.com/thanpawatpiti/notify/providers/telegram"
)

func main() {
	ctx := context.Background()
	msg := notify.Message{
		Title:    "Hello from Notify!",
		Content:  "This is a test notification sent from the Go notify library.",
		ImageURL: "https://via.placeholder.com/150",
		Color:    "#00FF00", // Green for Discord
	}

	// LINE
	if token := os.Getenv("LINE_CHANNEL_TOKEN"); token != "" {
		userID := os.Getenv("LINE_USER_ID")
		if userID != "" {
			log.Println("Sending LINE notification...")
			// Example of using Functional Options
			p := line.New(token, userID, notify.WithTimeout(10*time.Second))
			if err := p.Send(ctx, msg); err != nil {
				log.Printf("Failed to send LINE notification: %v", err)
			} else {
				log.Println("LINE notification sent!")
			}
		}
	}

	// Telegram
	if token := os.Getenv("TELEGRAM_TOKEN"); token != "" {
		chatID := os.Getenv("TELEGRAM_CHAT_ID")
		if chatID != "" {
			log.Println("Sending Telegram notification...")
			p := telegram.New(token, chatID)
			if err := p.Send(ctx, msg); err != nil {
				log.Printf("Failed to send Telegram notification: %v", err)
			} else {
				log.Println("Telegram notification sent!")
			}
		}
	}

	// Discord
	if webhookURL := os.Getenv("DISCORD_WEBHOOK_URL"); webhookURL != "" {
		log.Println("Sending Discord notification...")
		p := discord.New(webhookURL)
		if err := p.Send(ctx, msg); err != nil {
			log.Printf("Failed to send Discord notification: %v", err)
		} else {
			log.Println("Discord notification sent!")
		}
	}

	// MS Teams
	if webhookURL := os.Getenv("MSTEAMS_WEBHOOK_URL"); webhookURL != "" {
		log.Println("Sending MS Teams notification...")
		p := msteams.New(webhookURL)
		if err := p.Send(ctx, msg); err != nil {
			log.Printf("Failed to send MS Teams notification: %v", err)
		} else {
			log.Println("MS Teams notification sent!")
		}
	}
}
