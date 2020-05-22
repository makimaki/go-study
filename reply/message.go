package reply

import "encoding/json"

type TextMessage struct {
	Text string `json:"text"`
}

type TypedTextMessage TextMessage

func (m TextMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		TypedTextMessage
		Type string `json:"type"`
	}{
		TypedTextMessage: TypedTextMessage(m),
		Type:             "text",
	})
}

type TemplateMessage struct {
	AltText  string      `json:"altText"`
	Template interface{} `json:"template"`
}

type TypedTemplateMessage TemplateMessage

func (m TemplateMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		TypedTemplateMessage
		Type string `json:"type"`
	}{
		TypedTemplateMessage: TypedTemplateMessage(m),
		Type:                 "template",
	})
}
