package model

type Message struct {
	Model

	UserID       uint   `json:"userID"`
	ConnectionID uint   `json:"connectionID"`
	Content      string `json:"content"`

	User       *User       `json:"user"`
	Connection *Connection `json:"connection"`
}
