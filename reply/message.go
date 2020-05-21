package reply

type TextMessage struct {
	Text string `json:"text"`
}

type TemplateMessage struct {
	AltText  string      `json:"altText"`
	Template interface{} `json:"template"`
}
