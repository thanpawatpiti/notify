package line

// FlexMessage represents a LINE Flex Message.
type FlexMessage struct {
	AltText  string
	Contents FlexContainer
}

// FlexContainer is the interface for Flex Message containers (Bubble, Carousel).
type FlexContainer interface {
	isFlexContainer()
}

// BubbleContainer represents a Bubble container.
type BubbleContainer struct {
	Type   string          `json:"type"` // "bubble"
	Header *BoxComponent   `json:"header,omitempty"`
	Hero   *ImageComponent `json:"hero,omitempty"`
	Body   *BoxComponent   `json:"body,omitempty"`
	Footer *BoxComponent   `json:"footer,omitempty"`
	Styles *BubbleStyles   `json:"styles,omitempty"`
}

func (c BubbleContainer) isFlexContainer() {}

// CarouselContainer represents a Carousel container.
type CarouselContainer struct {
	Type     string            `json:"type"` // "carousel"
	Contents []BubbleContainer `json:"contents"`
}

func (c CarouselContainer) isFlexContainer() {}

// FlexComponent is the interface for Flex Message components.
type FlexComponent interface {
	isFlexComponent()
}

// BoxComponent represents a Box component.
type BoxComponent struct {
	Type     string          `json:"type"`   // "box"
	Layout   string          `json:"layout"` // "horizontal", "vertical", "baseline"
	Contents []FlexComponent `json:"contents"`
	Flex     *int            `json:"flex,omitempty"`
	Spacing  string          `json:"spacing,omitempty"`
	Margin   string          `json:"margin,omitempty"`
	Action   *Action         `json:"action,omitempty"`
}

func (c BoxComponent) isFlexComponent() {}

// TextComponent represents a Text component.
type TextComponent struct {
	Type   string  `json:"type"` // "text"
	Text   string  `json:"text"`
	Flex   *int    `json:"flex,omitempty"`
	Margin string  `json:"margin,omitempty"`
	Size   string  `json:"size,omitempty"`
	Align  string  `json:"align,omitempty"`
	Weight string  `json:"weight,omitempty"`
	Color  string  `json:"color,omitempty"`
	Wrap   bool    `json:"wrap,omitempty"`
	Action *Action `json:"action,omitempty"`
}

func (c TextComponent) isFlexComponent() {}

// ImageComponent represents an Image component.
type ImageComponent struct {
	Type        string  `json:"type"` // "image"
	URL         string  `json:"url"`
	Flex        *int    `json:"flex,omitempty"`
	Margin      string  `json:"margin,omitempty"`
	Align       string  `json:"align,omitempty"`
	Gravity     string  `json:"gravity,omitempty"`
	Size        string  `json:"size,omitempty"`
	AspectRatio string  `json:"aspectRatio,omitempty"`
	AspectMode  string  `json:"aspectMode,omitempty"`
	Action      *Action `json:"action,omitempty"`
}

func (c ImageComponent) isFlexComponent() {}

// ButtonComponent represents a Button component.
type ButtonComponent struct {
	Type   string `json:"type"` // "button"
	Action Action `json:"action"`
	Flex   *int   `json:"flex,omitempty"`
	Margin string `json:"margin,omitempty"`
	Height string `json:"height,omitempty"`
	Style  string `json:"style,omitempty"` // "link", "primary", "secondary"
	Color  string `json:"color,omitempty"`
}

func (c ButtonComponent) isFlexComponent() {}

// SeparatorComponent represents a Separator component.
type SeparatorComponent struct {
	Type   string `json:"type"` // "separator"
	Margin string `json:"margin,omitempty"`
	Color  string `json:"color,omitempty"`
}

func (c SeparatorComponent) isFlexComponent() {}

// Action represents an action.
type Action struct {
	Type  string `json:"type"`
	Label string `json:"label,omitempty"`
	URI   string `json:"uri,omitempty"`
	Data  string `json:"data,omitempty"`
	Text  string `json:"text,omitempty"`
}

// BubbleStyles represents styles for a Bubble container.
type BubbleStyles struct {
	Header *BlockStyle `json:"header,omitempty"`
	Hero   *BlockStyle `json:"hero,omitempty"`
	Body   *BlockStyle `json:"body,omitempty"`
	Footer *BlockStyle `json:"footer,omitempty"`
}

// BlockStyle represents style for a block.
type BlockStyle struct {
	BackgroundColor string `json:"backgroundColor,omitempty"`
	Separator       bool   `json:"separator,omitempty"`
	SeparatorColor  string `json:"separatorColor,omitempty"`
}
