package notify

import (
	"context"
	"net/http"
	"time"
)

// Notifier is the interface that all notification providers must implement.
type Notifier interface {
	// Send sends a payload to the provider.
	// payload can be:
	// - string: Simple text message.
	// - notify.CommonMessage: Generic rich message (Text + Image).
	// - Provider-specific structs: For advanced features (e.g., line.FlexMessage, discord.Embed).
	Send(ctx context.Context, payload interface{}) error
}

// CommonMessage represents a generic rich message supported by most providers.
// Use this for simple cross-platform notifications.
type CommonMessage struct {
	// Title is the subject or title of the message.
	Title string
	// Content is the main body of the message.
	Content string
	// ImageURL is an optional URL to an image.
	ImageURL string
	// Color is the color of the embed/message (Hex string e.g. "#FF0000").
	Color string
}

// Message is an alias for CommonMessage for backward compatibility (optional, but good for transition).
type Message = CommonMessage

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
func WithTimeout(d time.Duration) Option {
	return func(o *Options) {
		if o.HTTPClient == nil {
			o.HTTPClient = &http.Client{}
		}
		o.HTTPClient.Timeout = d
	}
}
