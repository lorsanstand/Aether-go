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

func fromUserRequest(RequestUser RequestUserUpdate) gen.UpdateUserParams {
	return gen.UpdateUserParams{
		DisplayName: RequestUser.DisplayName,
		Description: RequestUser.Description,
		BirthDay:    toPgDate(RequestUser.BirthDay),
		Username:    RequestUser.Username,
	}
}

func toPgDate(time2 time.Time) pgtype.Date {
	return pgtype.Date{
		Time:             time2,
		InfinityModifier: 0,
		Valid:            true,
	}
}
