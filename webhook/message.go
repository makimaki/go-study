package webhook

type Message struct {
	Id string
}

type TextMessage struct {
	Message
	Text string
}

type LocationMessage struct {
	Message
	Title     string `json:",omitempty"`
	Address   string `json:",omitempty"`
	Latitude  float64
	Longitude float64
}
