package discord

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/thanpawatpiti/notify"
)

// Provider implements the Notifier interface for Discord.
type Provider struct {
	webhookURL string
	opts       notify.Options
}

// New creates a new Discord provider.
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

// Send sends a message via Discord Webhook.
func (p *Provider) Send(ctx context.Context, msg notify.Message) error {
	if p.webhookURL == "" {
		return fmt.Errorf("discord webhook url is missing")
	}

	payload := map[string]interface{}{}

	// Create an embed for "beautiful" notification
	embed := map[string]interface{}{
		"description": msg.Content,
	}

	if msg.Title != "" {
		embed["title"] = msg.Title
	}

	if msg.ImageURL != "" {
		embed["image"] = map[string]string{
			"url": msg.ImageURL,
		}
	}

	if msg.Color != "" {
		if colorInt, err := parseColor(msg.Color); err == nil {
			embed["color"] = colorInt
		}
	}

	payload["embeds"] = []interface{}{embed}

	body, err := json.Marshal(payload)
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

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("discord webhook returned status: %d", resp.StatusCode)
	}

	return nil
}

func parseColor(colorStr string) (int, error) {
	colorStr = strings.TrimPrefix(colorStr, "#")
	val, err := strconv.ParseInt(colorStr, 16, 64)
	if err != nil {
		return 0, err
	}
	return int(val), nil
}
