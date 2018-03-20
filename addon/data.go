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

// icons
const (
	IconGithub   = "\uf09b"
	IconDisk     = "\uf1c0"
	IconCPU      = "\uf0e4"
	IconIP       = "\uf0e8"
	IconMemory   = "\uf2db"
	IconNetwork  = "\uf0c1"
	IconVolume   = "\uf028"
	IconTime     = "\uf017"
	IconPomodoro = "\uf0ae"
	IconWork     = "\uf0e7"
	IconPlay     = "\uf439"
)

const (
	ColorWhite   = "#FFFFFF"
	ColorSilver  = "#C0C0C0"
	ColorGray    = "#808080"
	ColorBlack   = "#000000"
	ColorRed     = "#FF0000"
	ColorMaroon  = "#800000"
	ColorYellow  = "#FFFF00"
	ColorOlive   = "#808000"
	ColorLime    = "#00FF00"
	ColorGreen   = "#008000"
	ColorAqua    = "#00FFFF"
	ColorTeal    = "#008080"
	ColorBlue    = "#0000FF"
	ColorNavy    = "#000080"
	ColorFuchsia = "#FF00FF"
	ColorPurple  = "#800080"
)
