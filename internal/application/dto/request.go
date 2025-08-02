package dto

type RequestData struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	User    string `json:"user"`
	Content string `json:"content"`
}
