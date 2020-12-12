package model

type FavoriteArtist struct {
	Model

	UserID      uint   `json:"userID"`
	FirstPlace  string `json:"firstPlace"`
	SocondPlace string `json:"socondPlace"`
	ThirdPlace  string `json:"thirdPlace"`
	FourthPlace string `json:"fourthPlace"`
	FifthPlace  string `json:"fifthPlace"`

	User *User `json:"user"`
}
