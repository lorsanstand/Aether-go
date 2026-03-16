package users

import "time"

type userDTO struct {
	Id          int32     `json:"id"`
	DisplayName string    `json:"display_name"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	BirthDay    time.Time `json:"birth_day"`
	Description *string   `json:"description"`
	AvatarUrl   *string   `json:"avatar_url"`
	IsActive    bool      `json:"is_active"`
	IsVerified  bool      `json:"is_verified"`
	IsSuperuser bool      `json:"is_superuser"`
}
