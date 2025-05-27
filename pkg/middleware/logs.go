package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		mw := &WrapperWriter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}
		next.ServeHTTP(mw, r)
		log.Printf("[Middleware] - [Logging] - [INFO] %d %s %s %s", mw.StatusCode, r.Method, r.URL.Path, time.Since(start))
	})
}
