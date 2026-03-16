package users

import (
	"net/http"
)

func (u *UserHandler) RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /users/me", u.authUsers(u.Me))
	mux.HandleFunc("GET /users/{id}", u.authUsers(u.GetUserById))

	return mux
}
