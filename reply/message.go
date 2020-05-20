package reply

type Message struct{}

type TextMessage struct {
	Message
	Text string `json:"text"`
}

type TemplateMessage struct {
	Message
	AltText  string      `json:"altText"`
	Template interface{} `json:"template"`
}
