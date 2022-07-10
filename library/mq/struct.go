package mq

type Email struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type Item struct {
	Retry int    `json:"retry"`
	Data  string `json:"data"`
}
