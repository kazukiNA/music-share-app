package model

type FavoriteGenre struct {
	Model

	UserID uint   `json:"userID"`
	Name   string `json:"name"`

	User *User `json:"user"`
}
