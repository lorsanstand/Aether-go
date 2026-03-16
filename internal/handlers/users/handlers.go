package users

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/lorsanstand/Aether-go/internal/database/sqlc/gen"
	"github.com/lorsanstand/Aether-go/pkg/respond"
)

type UserHandler struct {
	queries *gen.Queries
	respond.Respond
}

type User struct {
	User string `json:"user"`
	Pass string `json:"Password"`
}

func NewUserHandler(queries *gen.Queries) *UserHandler {
	return &UserHandler{queries: queries}
}

func (u *UserHandler) Me(w http.ResponseWriter, r *http.Request) {
	user := userFromContext(r.Context())
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
