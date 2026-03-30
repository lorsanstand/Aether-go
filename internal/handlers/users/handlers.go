package users

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/lorsanstand/Aether-go/internal/database/sqlc/gen"
	"github.com/lorsanstand/Aether-go/pkg/respond"
)

type UserStore interface {
	GetUserById(ctx context.Context, id int32) (gen.User, error)
	CreateUser(ctx context.Context, arg gen.CreateUserParams) error
	UpdateUser(ctx context.Context, arg gen.UpdateUserParams) error
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
			slog.InfoContext(r.Context(), "User not found")
		}
		return
	}

	user := fromDatabaseUsers(userDB)
	u.RespondJSON(w, http.StatusOK, user)
}

func (u *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := userFromContext(r.Context())
	if user.Username == "" {
		http.Error(w, "User unauthorized", http.StatusUnauthorized)
	}

	var RequestData RequestUserUpdate

	err := json.NewDecoder(r.Body).Decode(&RequestData)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		slog.WarnContext(r.Context(), "failed json validate", "error", err)
		return
	}

	err = u.queries.UpdateUser(r.Context(), fromUserRequest(RequestData))
	if err != nil {
		http.Error(w, "Failed save in DB", http.StatusInternalServerError)
		slog.ErrorContext(r.Context(), "failed save in DB", "error", err)
		return
	}
}
