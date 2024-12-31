package main

import (
	"fmt"
	"log"
	"net/http"

	"a1_ecominventory/database"
	"a1_ecominventory/handlers"
	"a1_ecominventory/middleware"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken() (string, error) {
	claims := jwt.MapClaims{
		"sub": "Venkatesh",
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("admin"))
}

func main() {
	token, err := GenerateToken()
	if err != nil {
		fmt.Println("Error generating token:", err)
		return
	}
	fmt.Println("Generated Token(Secret key to be entered in api):", token)

	database.Initialize()

	mux := http.NewServeMux()

	withMiddleware := func(handler http.Handler) http.Handler {
		return middleware.LoggingMiddleware(
			middleware.RateLimitMiddleware(
				middleware.AuthMiddleware(handler),
			),
		)
	}

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	mux.Handle("/product", withMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handlers.CreateProduct(w, r)
		case http.MethodGet:
			handlers.GetProducts(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})))

	mux.Handle("/product/", withMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetProduct(w, r)
		case http.MethodPut:
			handlers.UpdateProduct(w, r)
		case http.MethodDelete:
			handlers.DeleteProduct(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})))

	log.Printf("Server starting on http://localhost%s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
