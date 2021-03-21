package middlewares

import (
	"log"
	"net/http"
	"os"
	"time"
)

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if os.Getenv("ENV") != "production" {
			now := time.Now()
			next.ServeHTTP(w, r)
			since := time.Since(now).Milliseconds
			log.Printf("%v %v -- %v", r.Method, r.URL.Path, since())
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
