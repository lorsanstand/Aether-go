package users

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/lorsanstand/Aether-go/internal/database/sqlc/gen"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type userStoreMock struct {
}

func (u *userStoreMock) GetUserById(ctx context.Context, id int32) (gen.User, error) {
	if id == 1 {
		return gen.User{
			ID:             1,
			DisplayName:    "Test",
			Username:       "Test",
			Email:          "Test@example.com",
			BirthDay:       pgtype.Date{Time: time.Now(), Valid: true, InfinityModifier: pgtype.InfinityModifier(4)},
			Description:    nil,
			AvatarUrl:      nil,
			IsActive:       true,
			IsVerified:     true,
			IsSuperuser:    false,
			HashedPassword: "dsfsdfsdfs",
			CreatedAt:      pgtype.Timestamp{Time: time.Now(), Valid: true, InfinityModifier: pgtype.InfinityModifier(4)},
			UpdatedAt:      pgtype.Timestamp{Time: time.Now(), Valid: true, InfinityModifier: pgtype.InfinityModifier(4)},
		}, nil
	}
	return gen.User{}, sql.ErrNoRows
}

func TestUserHandler_GetUserById(t *testing.T) {
	mock := &userStoreMock{}
	handler := NewUserHandler(mock)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/users/1", nil)

	r.SetPathValue("id", "1")

	handler.GetUserById(w, r)

	assert.Equal(t, 200, w.Code, "Invalid status code")

	var user userDTO
	err := json.Unmarshal(w.Body.Bytes(), &user)
	require.NoError(t, err)
}

func TestUserHandler_GetUserById_UserNotFound(t *testing.T) {
	mock := &userStoreMock{}
	handler := NewUserHandler(mock)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/users/900", nil)

	r.SetPathValue("id", "900")

	handler.GetUserById(w, r)

	assert.Equal(t, 404, w.Code, "Invalid status code")
}

func TestUserHandler_GetUserById_BadUser(t *testing.T) {
	mock := &userStoreMock{}
	handler := NewUserHandler(mock)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/users/bad", nil)

	r.SetPathValue("id", "bad")

	handler.GetUserById(w, r)

	assert.Equal(t, 400, w.Code, "Invalid status code")
}

func TestUserHandler_Me(t *testing.T) {
	mock := &userStoreMock{}
	handler := NewUserHandler(mock)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/users/me/", nil)

	userDB, _ := mock.GetUserById(context.TODO(), 1)

	user := fromDatabaseUsers(userDB)

	r = r.WithContext(context.WithValue(r.Context(), userCtx, user))

	handler.Me(w, r)

	assert.Equal(t, 200, w.Code)

	var usr userDTO
	err := json.Unmarshal(w.Body.Bytes(), &usr)
	require.NoError(t, err)

	assert.NotEmpty(t, usr.Id, usr)
}

func TestUserHandler_Me_NoUser(t *testing.T) {
	mock := &userStoreMock{}
	handler := NewUserHandler(mock)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/users/me/", nil)

	handler.Me(w, r)

	require.Equal(t, 401, w.Code)

	//var usr userDTO
	//err := json.Unmarshal(w.Body.Bytes(), &usr)
	//require.NoError(t, err, "Failed marshal ")
	//
	//assert.Empty(t, usr.Username)
}
