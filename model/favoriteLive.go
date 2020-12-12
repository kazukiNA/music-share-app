package model

import "time"

type FavoriteLive struct {
	Model

	UserID     uint      `json:"userID"`
	ArtistName string    `json:"artistName"`
	EventDate  time.Time `json:"eventDate"`
	LiveName   string    `json:"liveName"`

	User *User `json:"user"`
}
