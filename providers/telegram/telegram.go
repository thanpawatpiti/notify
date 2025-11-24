package telegram

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/thanpawatpiti/notify"
)

const telegramAPIBase = "https://api.telegram.org/bot"

// Provider implements the Notifier interface for Telegram.
type Provider struct {
	token  string
	chatID string
	opts   notify.Options
}

// New creates a new Telegram provider.
func New(token, chatID string, opts ...notify.Option) *Provider {
	p := &Provider{
		token:  token,
		chatID: chatID,
		opts: notify.Options{
			HTTPClient: &http.Client{},
		},
	}

	for _, opt := range opts {
		opt(&p.opts)
	}

	return p
}

// Send sends a message via Telegram.
func (p *Provider) Send(ctx context.Context, msg notify.Message) error {
	if p.token == "" || p.chatID == "" {
		return fmt.Errorf("telegram token or chatID is missing")
	}

	var method string
	var payload interface{}

	text := msg.Content
	if msg.Title != "" {
		text = fmt.Sprintf("*%s*\n%s", msg.Title, msg.Content)
	}

	if msg.ImageURL != "" {
		method = "sendPhoto"
		payload = map[string]interface{}{
			"chat_id":    p.chatID,
			"photo":      msg.ImageURL,
			"caption":    text,
			"parse_mode": "Markdown",
		}
	} else {
		method = "sendMessage"
		payload = map[string]interface{}{
			"chat_id":    p.chatID,
			"text":       text,
			"parse_mode": "Markdown",
		}
	}

	url := fmt.Sprintf("%s%s/%s", telegramAPIBase, p.token, method)

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := p.opts.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("telegram api returned status: %d", resp.StatusCode)
	}

	return nil
}
