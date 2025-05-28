package middleware

import (
	"app/test/configs"
	"app/test/pkg/jwt"
	"context"
	"net/http"
	"strings"
)

type emailKey string

const (
	ContextEmailKey emailKey = "ContextEmailKey"
)

func writeUnauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}

func IsAuthed(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")

		if !strings.HasPrefix(authorizationHeader, "Bearer ") {
			writeUnauthorized(w)
			return
		}

		token := strings.TrimPrefix(authorizationHeader, "Bearer ")
		isValid, data := jwt.NewJWT(config.Auth.Secret).Parse(token)

		if !isValid {
			writeUnauthorized(w)
			return
		}

		ctx := context.WithValue(r.Context(), ContextEmailKey, data.Email)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
