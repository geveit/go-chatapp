package server

type Message struct {
	Sender  string `json:"user"`
	Content string `json:"message"`
}
