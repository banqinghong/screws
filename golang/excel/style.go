package excel

type FormatFont struct {
	Bold      bool   `json:"bold"`
	Italic    bool   `json:"italic"`
	Underline string `json:"underline"`
	Family    string `json:"family"`
	Size      int    `json:"size"`
	Color     string `json:"color"`
}

// formatStyle directly maps the styles settings of the cells.
type FormatStyle struct {
	Border []struct {
		Type  string `json:"type"`
		Color string `json:"color"`
		Style int    `json:"style"`
	} `json:"border"`
	Fill struct {
		Type    string   `json:"type"`
		Pattern int      `json:"pattern"`
		Color   []string `json:"color"`
		Shading int      `json:"shading"`
	} `json:"fill"`
	Font      *FormatFont `json:"font"`
	Alignment *struct {
		Horizontal      string `json:"horizontal"`
		Indent          int    `json:"indent"`
		JustifyLastLine bool   `json:"justify_last_line"`
		ReadingOrder    uint64 `json:"reading_order"`
		RelativeIndent  int    `json:"relative_indent"`
		ShrinkToFit     bool   `json:"shrink_to_fit"`
		TextRotation    int    `json:"text_rotation"`
		Vertical        string `json:"vertical"`
		WrapText        bool   `json:"wrap_text"`
	} `json:"alignment"`
	Protection *struct {
		Hidden bool `json:"hidden"`
		Locked bool `json:"locked"`
	} `json:"protection"`
	NumFmt        int     `json:"number_format"`
	DecimalPlaces int     `json:"decimal_places"`
	CustomNumFmt  *string `json:"custom_number_format"`
	Lang          string  `json:"lang"`
	NegRed        bool    `json:"negred"`
}
