package notify

import (
	"context"
	"net/http"
	"time"
)

// Notifier is the interface that all notification providers must implement.
type Notifier interface {
	// Send sends a message to the provider.
	Send(ctx context.Context, msg Message) error
}

// Message represents the content of the notification.
type Message struct {
	// Title is the subject or title of the message (optional, supported by some providers like Discord/Telegram/Teams).
	Title string
	// Content is the main body of the message.
	Content string
	// ImageURL is an optional URL to an image to include in the notification.
	ImageURL string
	// Color is the color of the embed/message (optional, supported by Discord/Teams).
	// Format: Hex string e.g. "#FF0000" or integer value.
	Color string
}

// Options holds common configuration for providers.
type Options struct {
	HTTPClient *http.Client
}

// Option is a function that configures Options.
type Option func(*Options)

// WithHTTPClient configures the provider to use a custom HTTP client.
func WithHTTPClient(client *http.Client) Option {
	return func(o *Options) {
		o.HTTPClient = client
	}
}

// WithTimeout configures a default timeout for the HTTP client if one isn't already set.
// Note: This is a convenience helper; usually it's better to set timeout on the client or context.
func WithTimeout(d time.Duration) Option {
	return func(o *Options) {
		if o.HTTPClient == nil {
			o.HTTPClient = &http.Client{}
		}
		o.HTTPClient.Timeout = d
	}
}
