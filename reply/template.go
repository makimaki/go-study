package reply

import "encoding/json"

type ButtonsTemplate struct {
	Text    string        `json:"text"`
	Actions []interface{} `json:"actions"`
}

type TypedButtonsTemplate ButtonsTemplate

func (t ButtonsTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		TypedButtonsTemplate
		Type string `json:"type"`
	}{
		TypedButtonsTemplate: TypedButtonsTemplate(t),
		Type:                 "buttons",
	})
}
