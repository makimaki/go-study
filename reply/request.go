package reply

type Request struct {
	ReplyToken string        `json:"replyToken"`
	Messages   []interface{} `json:"messages"`
}
