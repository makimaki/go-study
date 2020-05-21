package reply

type PostbackAction struct {
	Label       string `json:"label"`
	Data        string `json:"data"`
	DisplayText string `json:"displayText,omitempty"`
}

type UriAction struct {
	Label string `json:"label"`
	Uri   string `json:"uri"`
}
