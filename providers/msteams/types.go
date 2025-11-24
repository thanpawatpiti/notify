package msteams

// AdaptiveCard represents an Adaptive Card.
type AdaptiveCard struct {
	Type    string        `json:"type"`    // "AdaptiveCard"
	Version string        `json:"version"` // "1.2"
	Body    []interface{} `json:"body"`
	Actions []interface{} `json:"actions,omitempty"`
	Schema  string        `json:"$schema,omitempty"`
}

// TextBlock represents a TextBlock element.
type TextBlock struct {
	Type   string `json:"type"` // "TextBlock"
	Text   string `json:"text"`
	Size   string `json:"size,omitempty"`
	Weight string `json:"weight,omitempty"`
	Color  string `json:"color,omitempty"`
	Wrap   bool   `json:"wrap,omitempty"`
}

// Image represents an Image element.
type Image struct {
	Type string `json:"type"` // "Image"
	URL  string `json:"url"`
	Size string `json:"size,omitempty"`
	Alt  string `json:"altText,omitempty"`
}

// FactSet represents a FactSet element.
type FactSet struct {
	Type  string `json:"type"` // "FactSet"
	Facts []Fact `json:"facts"`
}

// Fact represents a fact in a FactSet.
type Fact struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

// ActionOpenUrl represents an OpenUrl action.
type ActionOpenUrl struct {
	Type  string `json:"type"` // "Action.OpenUrl"
	Title string `json:"title"`
	URL   string `json:"url"`
}

// WebhookPayload represents the payload sent to MS Teams Webhook.
type WebhookPayload struct {
	Type        string       `json:"type"` // "message"
	Attachments []Attachment `json:"attachments"`
}

// Attachment represents an attachment in the webhook payload.
type Attachment struct {
	ContentType string       `json:"contentType"` // "application/vnd.microsoft.card.adaptive"
	Content     AdaptiveCard `json:"content"`
}
