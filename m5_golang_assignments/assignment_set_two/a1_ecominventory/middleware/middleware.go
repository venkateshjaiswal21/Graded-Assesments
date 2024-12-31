package middleware

import (
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	rateLimiter = make(map[string][]time.Time)
	rateMutex   sync.Mutex
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("admin"), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		rateMutex.Lock()

		now := time.Now()
		var recent []time.Time
		for _, t := range rateLimiter[ip] {
			if now.Sub(t) < time.Minute {
				recent = append(recent, t)
			}
		}
		rateLimiter[ip] = recent

		if len(rateLimiter[ip]) >= 100 {
			rateMutex.Unlock()
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		rateLimiter[ip] = append(rateLimiter[ip], now)
		rateMutex.Unlock()

		next.ServeHTTP(w, r)
	})
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Request: Method=%s, URL=%s, RemoteAddr=%s",
			r.Method, r.URL.Path, r.RemoteAddr)

		next.ServeHTTP(w, r)

		log.Printf("Response: Completed in %v", time.Since(start))
	})
}
