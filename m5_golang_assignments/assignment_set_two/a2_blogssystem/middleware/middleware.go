package middleware

import (
	"log"
	"net/http"
	"time"
)

func ActivityLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		beginTime := time.Now()

		log.Printf("Access: %s %s from %s",
			r.Method,
			r.URL.Path,
			r.RemoteAddr,
		)

		next.ServeHTTP(w, r)

		duration := time.Since(beginTime)
		log.Printf("Request handled in %v", duration)
	})
}

func JSONValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
