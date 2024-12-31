package handlers

import (
	"a2_blogssystem/database"
	"a2_blogssystem/models"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.BlogPost
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid post data", http.StatusBadRequest)
		return
	}

	currentTime := time.Now()
	query := `INSERT INTO articles 
        (headline, article_text, contributor, created_date, last_modified) 
        VALUES (?, ?, ?, ?, ?)`

	result, err := database.DBConnection.Exec(query,
		post.Headline, post.ArticleText, post.Contributor, currentTime, currentTime)
	if err != nil {
		http.Error(w, "Failed to save post", http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	post.PostID = int(id)
	json.NewEncoder(w).Encode(post)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	postID := r.URL.Query().Get("id")

	query := `SELECT post_id, headline, article_text, contributor, 
        created_date, last_modified FROM articles WHERE post_id = ?`

	var post models.BlogPost
	err := database.DBConnection.QueryRow(query, postID).Scan(
		&post.PostID, &post.Headline, &post.ArticleText,
		&post.Contributor, &post.CreatedDate, &post.LastModified)

	if err == sql.ErrNoRows {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(post)
}

func ListPosts(w http.ResponseWriter, r *http.Request) {
	query := `SELECT post_id, headline, article_text, contributor, 
        created_date, last_modified FROM articles ORDER BY created_date DESC`

	rows, err := database.DBConnection.Query(query)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var posts []models.BlogPost
	for rows.Next() {
		var post models.BlogPost
		if err := rows.Scan(&post.PostID, &post.Headline, &post.ArticleText,
			&post.Contributor, &post.CreatedDate, &post.LastModified); err != nil {
			http.Error(w, "Data parsing error", http.StatusInternalServerError)
			return
		}
		posts = append(posts, post)
	}

	json.NewEncoder(w).Encode(posts)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	postID := r.URL.Query().Get("id")

	var post models.BlogPost
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid post data", http.StatusBadRequest)
		return
	}

	query := `UPDATE articles SET 
        headline = ?, article_text = ?, contributor = ?, last_modified = ? 
        WHERE post_id = ?`

	result, err := database.DBConnection.Exec(query,
		post.Headline, post.ArticleText, post.Contributor, time.Now(), postID)
	if err != nil {
		http.Error(w, "Update failed", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	w.Write([]byte("Post updated successfully"))
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	postID := r.URL.Query().Get("id")

	query := "DELETE FROM articles WHERE post_id = ?"
	result, err := database.DBConnection.Exec(query, postID)
	if err != nil {
		http.Error(w, "Deletion failed", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	w.Write([]byte("Post deleted successfully"))
}
