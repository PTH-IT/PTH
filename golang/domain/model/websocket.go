package model

type Websocket struct {
	Type     string `json:"type"`
	Receiver string `json:"receiver"`
	Sender   string `json:"sender"`
	Detail   string `json:"detail"`
}
