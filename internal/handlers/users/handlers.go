package users

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/lorsanstand/Aether-go/internal/database/sqlc/gen"
	"github.com/lorsanstand/Aether-go/pkg/respond"
)

type UserStore interface {
	GetUserById(ctx context.Context, id int32) (gen.User, error)
}

type UserHandler struct {
	queries UserStore
	respond.Respond
}

func NewUserHandler(queries UserStore) *UserHandler {
	return &UserHandler{queries: queries}
}

func (u *UserHandler) Me(w http.ResponseWriter, r *http.Request) {
	user := userFromContext(r.Context())
	if user.Username == "" {
		http.Error(w, "User unauthorized", http.StatusUnauthorized)
	}
	u.RespondJSON(w, http.StatusOK, user)
}

func (u *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Bad user id", http.StatusBadRequest)
		return
	}

	userDB, err := u.queries.GetUserById(r.Context(), int32(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "User not found", http.StatusNotFound)
		}
		return
	}

	user := fromDatabaseUsers(userDB)
	u.RespondJSON(w, http.StatusOK, user)
}
