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

	// 1. Simple Text Message (Supported by all)
	simpleMsg := "Hello from Notify! This is a simple text message."

	// 2. Common Rich Message (Supported by all)
	commonMsg := notify.CommonMessage{
		Title:    "Rich Notification",
		Content:  "This is a rich message with title and image.",
		ImageURL: "https://via.placeholder.com/150",
		Color:    "#00FF00",
	}

	// LINE
	if token := os.Getenv("LINE_CHANNEL_TOKEN"); token != "" {
		userID := os.Getenv("LINE_USER_ID")
		if userID != "" {
			log.Println("Sending LINE notifications...")
			p := line.New(token, userID, notify.WithTimeout(10*time.Second))

			// Send Simple Text
			p.Send(ctx, simpleMsg)

			// Send Common Message
			p.Send(ctx, commonMsg)

			// Send Advanced Flex Message
			flexMsg := line.FlexMessage{
				AltText: "Flex Message Example",
				Contents: line.BubbleContainer{
					Type: "bubble",
					Body: &line.BoxComponent{
						Type:   "box",
						Layout: "vertical",
						Contents: []line.FlexComponent{
							line.TextComponent{
								Type:   "text",
								Text:   "Flex Message",
								Weight: "bold",
								Size:   "xl",
							},
							line.TextComponent{
								Type: "text",
								Text: "This is a complex Flex Message layout.",
								Wrap: true,
							},
						},
					},
				},
			}
			if err := p.Send(ctx, flexMsg); err != nil {
				log.Printf("Failed to send LINE Flex Message: %v", err)
			}
		}
	}

	// Telegram
	if token := os.Getenv("TELEGRAM_TOKEN"); token != "" {
		chatID := os.Getenv("TELEGRAM_CHAT_ID")
		if chatID != "" {
			log.Println("Sending Telegram notifications...")
			p := telegram.New(token, chatID)

			// Send Simple Text
			p.Send(ctx, simpleMsg)

			// Send Advanced Payload with Keyboard
			payload := telegram.Payload{
				Text:      "Check out this keyboard!",
				ParseMode: "Markdown",
				ReplyMarkup: telegram.InlineKeyboardMarkup{
					InlineKeyboard: [][]telegram.InlineKeyboardButton{
						{
							{Text: "Google", URL: "https://google.com"},
							{Text: "GitHub", URL: "https://github.com"},
						},
					},
				},
			}
			if err := p.Send(ctx, payload); err != nil {
				log.Printf("Failed to send Telegram advanced message: %v", err)
			}
		}
	}

	// Discord
	if webhookURL := os.Getenv("DISCORD_WEBHOOK_URL"); webhookURL != "" {
		log.Println("Sending Discord notifications...")
		p := discord.New(webhookURL)

		// Send Common Message
		p.Send(ctx, commonMsg)

		// Send Advanced Embed
		embed := discord.Embed{
			Title:       "Advanced Embed",
			Description: "This embed has fields and a footer.",
			Color:       0xFF0000,
			Fields: []discord.EmbedField{
				{Name: "Field 1", Value: "Value 1", Inline: true},
				{Name: "Field 2", Value: "Value 2", Inline: true},
			},
			Footer: &discord.EmbedFooter{
				Text: "Sent via Notify",
			},
		}
		if err := p.Send(ctx, embed); err != nil {
			log.Printf("Failed to send Discord embed: %v", err)
		}
	}

	// MS Teams
	if webhookURL := os.Getenv("MSTEAMS_WEBHOOK_URL"); webhookURL != "" {
		log.Println("Sending MS Teams notifications...")
		p := msteams.New(webhookURL)

		// Send Common Message
		p.Send(ctx, commonMsg)

		// Send Advanced Adaptive Card
		card := msteams.AdaptiveCard{
			Type:    "AdaptiveCard",
			Version: "1.2",
			Body: []interface{}{
				msteams.TextBlock{
					Type:   "TextBlock",
					Text:   "Adaptive Card Example",
					Size:   "Large",
					Weight: "Bolder",
				},
				msteams.FactSet{
					Type: "FactSet",
					Facts: []msteams.Fact{
						{Title: "Fact 1", Value: "Value 1"},
						{Title: "Fact 2", Value: "Value 2"},
					},
				},
			},
		}
		if err := p.Send(ctx, card); err != nil {
			log.Printf("Failed to send MS Teams card: %v", err)
		}
	}
}
