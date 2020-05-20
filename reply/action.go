package reply

type Action struct {
	Label string `json:"label"`
}

type PostbackAction struct {
	Action
	Data        string `json:"data"`
	DisplayText string `json:"displayText,omitempty"`
}

type UriAction struct {
	Action
	Uri string `json:"uri"`
}
