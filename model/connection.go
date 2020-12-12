package model

type Connection struct {
	Model

	PreviousUserID  uint   `json:"previousUserID"`
	FollowingUserID uint   `json:"followingUserID"`
	ConnectionID    uint   `json:"connectionID"`
	Status          string `json:"status"`

	User *User `json:"user"`
}
