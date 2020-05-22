package reply

import "encoding/json"

type PostbackAction struct {
	Label       string `json:"label"`
	Data        string `json:"data"`
	DisplayText string `json:"displayText,omitempty"`
}

type TypedPostbackAction PostbackAction

func (a PostbackAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		TypedPostbackAction
		Type string `json:"type"`
	}{
		TypedPostbackAction: TypedPostbackAction(a),
		Type:                "postback",
	})
}

type UriAction struct {
	Label string `json:"label"`
	Uri   string `json:"uri"`
}

type TypedUriAction UriAction

func (a UriAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		TypedUriAction
		Type string `json:"type"`
	}{
		TypedUriAction: TypedUriAction(a),
		Type:           "uri",
	})
}
