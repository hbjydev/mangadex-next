package middlewares

import (
	"log"
	"net/http"
	"time"
)

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		next.ServeHTTP(w, r)
		since := time.Since(now).Milliseconds
		log.Printf("%v %v -- %v", r.Method, r.URL.Path, since())
	})
}
