package line

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/thanpawatpiti/notify"
)

const lineMessagingAPI = "https://api.line.me/v2/bot/message/push"

// Provider implements the Notifier interface for LINE Messaging API.
type Provider struct {
	channelToken string
	targetID     string // UserID or GroupID
	opts         notify.Options
}

// New creates a new LINE Messaging API provider.
func New(channelToken, targetID string, opts ...notify.Option) *Provider {
	p := &Provider{
		channelToken: channelToken,
		targetID:     targetID,
		opts: notify.Options{
			HTTPClient: &http.Client{},
		},
	}

	for _, opt := range opts {
		opt(&p.opts)
	}

	return p
}

// Send sends a message via LINE Messaging API.
func (p *Provider) Send(ctx context.Context, msg notify.Message) error {
	if p.channelToken == "" || p.targetID == "" {
		return fmt.Errorf("line channel token or target ID is missing")
	}

	var messages []interface{}

	// If image is present, send image message
	if msg.ImageURL != "" {
		messages = append(messages, map[string]string{
			"type":               "image",
			"originalContentUrl": msg.ImageURL,
			"previewImageUrl":    msg.ImageURL,
		})
	}

	// Always send text content if present (or as caption if image exists, but LINE separates them)
	if msg.Content != "" {
		text := msg.Content
		if msg.Title != "" {
			text = fmt.Sprintf("%s\n%s", msg.Title, msg.Content)
		}
		messages = append(messages, map[string]string{
			"type": "text",
			"text": text,
		})
	}

	payload := map[string]interface{}{
		"to":       p.targetID,
		"messages": messages,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, lineMessagingAPI, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+p.channelToken)

	resp, err := p.opts.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("line messaging api returned status: %d", resp.StatusCode)
	}

	return nil
}
