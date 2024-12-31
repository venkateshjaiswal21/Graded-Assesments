package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Initialize() {
	var err error
	DB, err = sql.Open("sqlite", "./inventory.db")
	if err != nil {
		log.Fatalf("Failed to connect to SQLite: %v", err)
	}

	createProductsTable := `
	CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		price REAL NOT NULL,
		stock INTEGER NOT NULL,
		category_id INTEGER,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	createCategoriesTable := `
	CREATE TABLE IF NOT EXISTS categories (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL UNIQUE
	);`

	if _, err := DB.Exec(createProductsTable); err != nil {
		log.Fatalf("Failed to create products table: %v", err)
	}

	if _, err := DB.Exec(createCategoriesTable); err != nil {
		log.Fatalf("Failed to create categories table: %v", err)
	}

	log.Println("Database initialized successfully.")
}
