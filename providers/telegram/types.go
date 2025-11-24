package telegram

// Payload represents a Telegram message payload.
type Payload struct {
	ChatID                string      `json:"chat_id"`
	Text                  string      `json:"text,omitempty"`
	ParseMode             string      `json:"parse_mode,omitempty"` // "MarkdownV2", "HTML", "Markdown"
	DisableWebPagePreview bool        `json:"disable_web_page_preview,omitempty"`
	DisableNotification   bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID      int         `json:"reply_to_message_id,omitempty"`
	ReplyMarkup           interface{} `json:"reply_markup,omitempty"` // InlineKeyboardMarkup or ReplyKeyboardMarkup
	Photo                 string      `json:"photo,omitempty"`        // URL for sendPhoto
	Caption               string      `json:"caption,omitempty"`      // For sendPhoto
}

// InlineKeyboardMarkup represents an inline keyboard.
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

// InlineKeyboardButton represents a button in an inline keyboard.
type InlineKeyboardButton struct {
	Text         string `json:"text"`
	URL          string `json:"url,omitempty"`
	CallbackData string `json:"callback_data,omitempty"`
}

// ReplyKeyboardMarkup represents a custom keyboard.
type ReplyKeyboardMarkup struct {
	Keyboard        [][]KeyboardButton `json:"keyboard"`
	ResizeKeyboard  bool               `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard bool               `json:"one_time_keyboard,omitempty"`
}

// KeyboardButton represents a button in a custom keyboard.
type KeyboardButton struct {
	Text string `json:"text"`
}
