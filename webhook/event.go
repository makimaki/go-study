package webhook

import (
	"encoding/json"
	"log"
)

type MessageEvent struct {
	Timestamp  int64
	Source     interface{}
	ReplyToken string
	Message    interface{}
}

func (this *MessageEvent) UnmarshalJSON(b []byte) error {
	var unsafe struct {
		Timestamp  int64
		Source     interface{}
		ReplyToken string
		Message    json.RawMessage
	}

	if err := json.Unmarshal(b, &unsafe); err != nil {
		return err
	}

	this.ReplyToken = unsafe.ReplyToken
	this.Source = unsafe.ReplyToken
	this.Timestamp = unsafe.Timestamp

	var m map[string]interface{}
	if err := json.Unmarshal(unsafe.Message, &m); err != nil {
		return err
	}
	switch m["type"] {
	case "text":
		var text TextMessage
		if err := json.Unmarshal(unsafe.Message, &text); err != nil {
			return err
		}

		this.Message = text
	case "location":
		var location LocationMessage
		if err := json.Unmarshal(unsafe.Message, &location); err != nil {
			return err
		}

		this.Message = location
	default:
		log.Printf("unknown message: %#v\n", m)
	}

	return nil
}

type PostbackEvent struct {
	Timestamp  int64
	Source     interface{}
	ReplyToken string
	Postback   struct {
		Data string
	}
}
