package addon

type Block struct {
	FullText            string                 `json:"full_text"`
	ShortText           string                 `json:"short_text,omitempty"`
	Color               string                 `json:"color,omitempty"`
	BorderColor         string                 `json:"border,omitempty"`
	BackgroundColor     string                 `json:"background,omitempty"`
	Markup              string                 `json:"markup,omitempty"`
	MinWidth            string                 `json:"min_width,omitempty"`
	Align               string                 `json:"align,omitempty"`
	Name                string                 `json:"name,omitempty"`
	Instance            string                 `json:"instance,omitempty"`
	Urgent              bool                   `json:"urgent,omitempty"`
	Separator           *bool                  `json:"separator,omitempty"`
	SeparatorBlockWidth uint16                 `json:"separator_block_width,omitempty"`
	Custom              map[string]interface{} `json:"-"`
}

type ClickEvent struct {
	Name      string `json:"name,omitempty"`
	Instance  string `json:"instance,omitempty"`
	Button    uint8  `json:"button"`
	X         uint16 `json:"x"`
	Y         uint16 `json:"y"`
	RelativeX uint16 `json:"relative_x"`
	RelativeY uint16 `json:"relative_y"`
	Width     uint16 `json:"width"`
	Height    uint16 `json:"height"`
}
