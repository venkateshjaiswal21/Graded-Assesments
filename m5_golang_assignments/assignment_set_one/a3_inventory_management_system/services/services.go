package services

import (
	"a3_inventory_management_system/models"
	"errors"
	"fmt"
	"sort"
	"strings"
)

var inventory []models.Product

func AddProduct(id int, name string, price float64, stock int) error {
	if id <= 0 {
		return errors.New("product ID must be positive")
	}
	if price < 0 {
		return errors.New("price cannot be negative")
	}
	if stock < 0 {
		return errors.New("stock cannot be negative")
	}

	for _, product := range inventory {
		if product.ID == id {
			return errors.New("product with this ID already exists")
		}
	}

	inventory = append(inventory, models.Product{
		ID:    id,
		Name:  name,
		Price: price,
		Stock: stock,
	})
	return nil
}

func UpdateStock(id int, newStock int) error {
	if newStock < 0 {
		return errors.New("stock cannot be negative")
	}

	for i, product := range inventory {
		if product.ID == id {
			inventory[i].Stock = newStock
			return nil
		}
	}
	return errors.New("product not found")
}

func SearchProduct(query string) (models.Product, error) {
	query = strings.ToLower(query)

	for _, product := range inventory {
		if strings.ToLower(product.Name) == query || fmt.Sprintf("%d", product.ID) == query {
			return product, nil
		}
	}
	return models.Product{}, errors.New("product not found")
}

func GetInventory() []models.Product {
	return inventory
}

func SortInventory(by string) ([]models.Product, error) {
	switch strings.ToLower(by) {
	case "price":
		sort.Slice(inventory, func(i, j int) bool {
			return inventory[i].Price < inventory[j].Price
		})
	case "stock":
		sort.Slice(inventory, func(i, j int) bool {
			return inventory[i].Stock < inventory[j].Stock
		})
	default:
		return nil, errors.New("invalid sorting criteria, use 'price' or 'stock'")
	}

	return inventory, nil
}
