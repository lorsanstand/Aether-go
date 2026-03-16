package users

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/lorsanstand/Aether-go/internal/handlers/middlewares"
)

type ctxKey string

const UserCtx ctxKey = "User"

func (u *UserHandler) authUsers(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, ok := middlewares.UserIdFromContext(r.Context())
		if !ok {
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}

		userDB, err := u.queries.GetUserById(r.Context(), int32(userId))
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}
		}

		user := fromDatabaseUsers(userDB)

		r = r.WithContext(context.WithValue(r.Context(), UserCtx, user))

		next(w, r)
	}
}

func userFromContext(ctx context.Context) userDTO {
	user, ok := ctx.Value(UserCtx).(userDTO)
	if !ok {
		return userDTO{}
	}
	return user
}
