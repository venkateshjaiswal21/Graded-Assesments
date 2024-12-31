package main

import (
	"a2_blogssystem/database"
	"a2_blogssystem/handlers"
	"a2_blogssystem/middleware"
	"log"
	"net/http"
)

func main() {
	database.InitDB()

	mux := http.NewServeMux()

	withMiddleware := func(handler http.HandlerFunc) http.Handler {
		return middleware.ActivityLogger(
			middleware.JSONValidator(
				http.HandlerFunc(handler),
			),
		)
	}

	mux.Handle("/post", withMiddleware(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handlers.CreatePost(w, r)
		case http.MethodGet:
			handlers.ListPosts(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))

	mux.Handle("/post/", withMiddleware(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetPost(w, r)
		case http.MethodPut:
			handlers.UpdatePost(w, r)
		case http.MethodDelete:
			handlers.DeletePost(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))

	log.Printf("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
