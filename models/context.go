package models

type ContextMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Context struct {
	Model    string           `json:"model"`
	Messages []ContextMessage `json:"messages"`
}
