package reply

type ButtonTemplate struct {
	Text    string        `json:"text"`
	Actions []interface{} `json:"actions"`
}
