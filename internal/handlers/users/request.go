package users

import "time"

type RequestUserUpdate struct {
	DisplayName string    `json:"display_name"`
	Username    string    `json:"username"`
	BirthDay    time.Time `json:"birth_day"`
	Description *string   `json:"description"`
}
