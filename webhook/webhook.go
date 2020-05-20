package webhook

import (
	"encoding/json"
	"log"
)

type Request struct {
	Events []interface{}
}

func (this *Request) UnmarshalJSON(b []byte) error {
	var request struct {
		Events []json.RawMessage
	}
	if err := json.Unmarshal(b, &request); err != nil {
		return err
	}
	for _, event := range request.Events {
		var e map[string]interface{}
		if err := json.Unmarshal(event, &e); err != nil {
			return err
		}
		switch e["type"] {
		case "message":
			var messageEvent MessageEvent
			if err := json.Unmarshal(event, &messageEvent); err != nil {
				log.Fatal(err)
			}
			log.Println(messageEvent)

			this.Events = append(this.Events, messageEvent)
		case "postback":
			var postbackEvent PostbackEvent
			if err := json.Unmarshal(event, &postbackEvent); err != nil {
				log.Fatal(err)
			}
			log.Println(postbackEvent)

			this.Events = append(this.Events, postbackEvent)
		default:
			log.Printf("unknown event: %#v\n", event)
		}
	}
	return nil
}
