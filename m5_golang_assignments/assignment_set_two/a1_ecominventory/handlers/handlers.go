package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"a1_ecominventory/database"
	"a1_ecominventory/models"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if product.Name == "" || product.Price <= 0 || product.Stock < 0 {
		http.Error(w, "Invalid product data", http.StatusBadRequest)
		return
	}

	now := time.Now()
	query := `
		INSERT INTO products (name, description, price, stock, category_id, created_at, updated_at) 
		VALUES (?, ?, ?, ?, ?, ?, ?)`

	result, err := database.DB.Exec(query,
		product.Name,
		product.Description,
		product.Price,
		product.Stock,
		product.CategoryID,
		now,
		now,
	)
	if err != nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	product.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	query := `
		SELECT id, name, description, price, stock, category_id, created_at, updated_at 
		FROM products WHERE id = ?`

	var product models.Product
	err = database.DB.QueryRow(query, id).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Stock,
		&product.CategoryID,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Failed to fetch product", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	query := `
		SELECT id, name, description, price, stock, category_id, created_at, updated_at 
		FROM products 
		ORDER BY created_at DESC 
		LIMIT ? OFFSET ?`

	rows, err := database.DB.Query(query, limit, offset)
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Stock,
			&product.CategoryID,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			http.Error(w, "Error reading product data", http.StatusInternalServerError)
			return
		}
		products = append(products, product)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	query := `
		UPDATE products 
		SET name = ?, description = ?, price = ?, stock = ?, category_id = ?, updated_at = ?
		WHERE id = ?`

	_, err = database.DB.Exec(query,
		product.Name,
		product.Description,
		product.Price,
		product.Stock,
		product.CategoryID,
		time.Now(),
		id,
	)

	if err != nil {
		http.Error(w, "Failed to update product", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Product updated successfully"))
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	query := "DELETE FROM products WHERE id = ?"
	result, err := database.DB.Exec(query, id)
	if err != nil {
		http.Error(w, "Failed to delete product", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.Write([]byte("Product deleted successfully"))
}
