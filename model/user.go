package model

type User struct {
	Model

	Name           string `json:"name"`
	Age            int    `json:"age"`
	Profession     string `json:"pofession"`
	ProfileContent string `json:"profileContent"`
	ProfileImage   string `json:"profileImage"`
}
