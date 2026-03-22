package middlewares

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	slogctx "github.com/veqryn/slog-context"
)

const authPrefix = "Bearer "

type ctxKey string

const userIDCtx ctxKey = "UserId"

func GetUserIdMiddleware(next http.Handler, secret string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		userID, ok := getUserIdWithToken(w, r, secret)
		if !ok {
			next.ServeHTTP(w, r)
			return
		}

		ctx := r.Context()
		ctx = slogctx.Append(ctx, "user_id", userID)

		r = r.WithContext(ctx)
		r = r.WithContext(context.WithValue(r.Context(), userIDCtx, userID))

		next.ServeHTTP(w, r)
	})
}

func getUserIdWithToken(w http.ResponseWriter, r *http.Request, secret string) (int, bool) {
	tokenCookie, err := r.Cookie("access_token")

	if err != nil || !strings.HasPrefix(tokenCookie.Value, authPrefix) {
		slog.Debug("failed get token", "error", err)
		return 0, false
	}

	tokenStr := strings.TrimPrefix(tokenCookie.Value, authPrefix)

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		slog.Warn("invalid token", "error", err)
		return 0, false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		slog.Warn("failed token claims")
		return 0, false
	}

	val, ok := claims["sub"].(string)
	if !ok {
		slog.Warn("not found user id in token")
		return 0, false
	}
	id, err := strconv.Atoi(val)
	if err != nil {
		slog.Warn("bad user id int token")
		return 0, false
	}

	return id, true
}

func UserIdFromContext(ctx context.Context) (int, bool) {
	id, ok := ctx.Value(userIDCtx).(int)
	return id, ok
}
