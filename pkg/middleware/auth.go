package middleware

import (
	"log"
	"net/http"
	"strings"
)

func IsAuthed(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		token := strings.TrimPrefix(authorizationHeader, "Bearer ")
		log.Println(token)
		next.ServeHTTP(w, r)
	})
}
