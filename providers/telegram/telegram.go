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
// payload can be:
// - string: Simple text message.
// - notify.CommonMessage: Generic rich message (Text + Image).
// - telegram.Payload: Full API payload.
func (p *Provider) Send(ctx context.Context, payload interface{}) error {
	if p.token == "" || p.chatID == "" {
		return fmt.Errorf("telegram token or chatID is missing")
	}

	var method string = "sendMessage"
	var reqPayload Payload

	switch v := payload.(type) {
	case string:
		reqPayload = Payload{
			ChatID:    p.chatID,
			Text:      v,
			ParseMode: "Markdown",
		}
	case notify.CommonMessage:
		text := v.Content
		if v.Title != "" {
			text = fmt.Sprintf("*%s*\n%s", v.Title, v.Content)
		}
		if v.ImageURL != "" {
			method = "sendPhoto"
			reqPayload = Payload{
				ChatID:    p.chatID,
				Photo:     v.ImageURL,
				Caption:   text,
				ParseMode: "Markdown",
			}
		} else {
			reqPayload = Payload{
				ChatID:    p.chatID,
				Text:      text,
				ParseMode: "Markdown",
			}
		}
	case Payload:
		reqPayload = v
		if reqPayload.ChatID == "" {
			reqPayload.ChatID = p.chatID
		}
		if reqPayload.Photo != "" {
			method = "sendPhoto"
		}
	default:
		return fmt.Errorf("unsupported payload type: %T", v)
	}

	url := fmt.Sprintf("%s%s/%s", telegramAPIBase, p.token, method)

	body, err := json.Marshal(reqPayload)
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
