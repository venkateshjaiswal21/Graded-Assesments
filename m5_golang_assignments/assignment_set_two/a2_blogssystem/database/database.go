package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DBConnection *sql.DB

func InitDB() {
	var err error
	DBConnection, err = sql.Open("sqlite", "./articles.db")
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	createTable := `
    CREATE TABLE IF NOT EXISTS articles (
        post_id INTEGER PRIMARY KEY AUTOINCREMENT,
        headline TEXT NOT NULL,
        article_text TEXT NOT NULL,
        contributor TEXT NOT NULL,
        created_date DATETIME DEFAULT CURRENT_TIMESTAMP,
        last_modified DATETIME DEFAULT CURRENT_TIMESTAMP
    );`

	if _, err := DBConnection.Exec(createTable); err != nil {
		log.Fatalf("Table creation failed: %v", err)
	}
}
