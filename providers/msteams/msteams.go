package msteams

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/thanpawatpiti/notify"
)

// Provider implements the Notifier interface for Microsoft Teams.
type Provider struct {
	webhookURL string
	opts       notify.Options
}

// New creates a new Microsoft Teams provider.
func New(webhookURL string, opts ...notify.Option) *Provider {
	p := &Provider{
		webhookURL: webhookURL,
		opts: notify.Options{
			HTTPClient: &http.Client{},
		},
	}

	for _, opt := range opts {
		opt(&p.opts)
	}

	return p
}

// Send sends a message via Microsoft Teams Incoming Webhook.
// payload can be:
// - string: Simple text message.
// - notify.CommonMessage: Generic rich message (Text + Image).
// - msteams.AdaptiveCard: Full Adaptive Card.
func (p *Provider) Send(ctx context.Context, payload interface{}) error {
	if p.webhookURL == "" {
		return fmt.Errorf("msteams webhook url is missing")
	}

	var card AdaptiveCard

	switch v := payload.(type) {
	case string:
		card = AdaptiveCard{
			Type:    "AdaptiveCard",
			Version: "1.2",
			Schema:  "http://adaptivecards.io/schemas/adaptive-card.json",
			Body: []interface{}{
				TextBlock{
					Type: "TextBlock",
					Text: v,
					Wrap: true,
				},
			},
		}
	case notify.CommonMessage:
		body := []interface{}{}
		if v.Title != "" {
			body = append(body, TextBlock{
				Type:   "TextBlock",
				Text:   v.Title,
				Weight: "Bolder",
				Size:   "Medium",
			})
		}
		if v.Content != "" {
			body = append(body, TextBlock{
				Type: "TextBlock",
				Text: v.Content,
				Wrap: true,
			})
		}
		if v.ImageURL != "" {
			body = append(body, Image{
				Type: "Image",
				URL:  v.ImageURL,
				Size: "Stretch",
			})
		}
		card = AdaptiveCard{
			Type:    "AdaptiveCard",
			Version: "1.2",
			Schema:  "http://adaptivecards.io/schemas/adaptive-card.json",
			Body:    body,
		}
	case AdaptiveCard:
		card = v
		if card.Schema == "" {
			card.Schema = "http://adaptivecards.io/schemas/adaptive-card.json"
		}
		if card.Type == "" {
			card.Type = "AdaptiveCard"
		}
		if card.Version == "" {
			card.Version = "1.2"
		}
	default:
		return fmt.Errorf("unsupported payload type: %T", v)
	}

	wp := WebhookPayload{
		Type: "message",
		Attachments: []Attachment{
			{
				ContentType: "application/vnd.microsoft.card.adaptive",
				Content:     card,
			},
		},
	}

	body, err := json.Marshal(wp)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, p.webhookURL, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := p.opts.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted {
		return fmt.Errorf("msteams webhook returned status: %d", resp.StatusCode)
	}

	return nil
}
