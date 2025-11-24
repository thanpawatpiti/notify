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

// Send sends a message via Microsoft Teams Incoming Webhook using Adaptive Cards.
func (p *Provider) Send(ctx context.Context, msg notify.Message) error {
	if p.webhookURL == "" {
		return fmt.Errorf("msteams webhook url is missing")
	}

	// Construct Adaptive Card
	card := map[string]interface{}{
		"type": "message",
		"attachments": []interface{}{
			map[string]interface{}{
				"contentType": "application/vnd.microsoft.card.adaptive",
				"content": map[string]interface{}{
					"$schema": "http://adaptivecards.io/schemas/adaptive-card.json",
					"type":    "AdaptiveCard",
					"version": "1.2",
					"body":    []interface{}{},
				},
			},
		},
	}

	bodyContent := []interface{}{}

	// Title
	if msg.Title != "" {
		bodyContent = append(bodyContent, map[string]interface{}{
			"type":   "TextBlock",
			"text":   msg.Title,
			"weight": "Bolder",
			"size":   "Medium",
		})
	}

	// Content
	if msg.Content != "" {
		bodyContent = append(bodyContent, map[string]interface{}{
			"type": "TextBlock",
			"text": msg.Content,
			"wrap": true,
		})
	}

	// Image
	if msg.ImageURL != "" {
		bodyContent = append(bodyContent, map[string]interface{}{
			"type": "Image",
			"url":  msg.ImageURL,
			"size": "Stretch",
		})
	}

	// Assign body to card
	card["attachments"].([]interface{})[0].(map[string]interface{})["content"].(map[string]interface{})["body"] = bodyContent

	body, err := json.Marshal(card)
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
