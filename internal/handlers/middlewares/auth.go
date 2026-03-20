package middlewares

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

const authPrefix = "Bearer "

type ctxKey string

const userIdCtx ctxKey = "UserId"

func GetUserIdMiddleware(next http.Handler, secret string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		userId, ok := getUserIdWithToken(w, r, secret)
		if !ok {
			next.ServeHTTP(w, r)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), userIdCtx, userId))

		next.ServeHTTP(w, r)
	})
}

func getUserIdWithToken(w http.ResponseWriter, r *http.Request, secret string) (int, bool) {
	tokenCookie, err := r.Cookie("access_token")

	if err != nil || !strings.HasPrefix(tokenCookie.Value, authPrefix) {
		log.Printf("Failed get token: %v", err)
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
		log.Printf("Invalid token: %v", err)
		return 0, false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Printf("Failed token claims")
		return 0, false
	}

	val, ok := claims["sub"].(string)
	if !ok {
		log.Printf("Invalid token not found user id")
		return 0, false
	}
	id, err := strconv.Atoi(val)
	if err != nil {
		log.Printf("Bad user id")
		return 0, false
	}

	return id, true
}

func UserIdFromContext(ctx context.Context) (int, bool) {
	id, ok := ctx.Value(userIdCtx).(int)
	return id, ok
}
