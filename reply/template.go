package reply

type ButtonsTemplate struct {
	Text    string        `json:"text"`
	Actions []interface{} `json:"actions"`
}
