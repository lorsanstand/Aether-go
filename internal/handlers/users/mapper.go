package users

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/lorsanstand/Aether-go/internal/database/sqlc/gen"
)

func fromDatabaseUsers(user gen.User) userDTO {
	return userDTO{
		Id:          user.ID,
		DisplayName: user.DisplayName,
		Username:    user.Username,
		Email:       user.Email,
		BirthDay:    toTime(user.BirthDay),
		Description: user.Description,
		AvatarUrl:   user.AvatarUrl,
		IsActive:    user.IsActive,
		IsVerified:  user.IsVerified,
		IsSuperuser: user.IsSuperuser,
	}
}

func toTime(pgDate pgtype.Date) time.Time {
	if !pgDate.Valid {
		return time.Time{}
	}
	return pgDate.Time
}
