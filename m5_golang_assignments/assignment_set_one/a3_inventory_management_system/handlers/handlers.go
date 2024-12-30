package handlers

import (
	"a3_inventory_management_system/services"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func AddProductHandler(scanner *bufio.Scanner) {
	fmt.Print("Enter Product ID: ")
	scanner.Scan()
	id, err := strconv.Atoi(scanner.Text())
	if err != nil || id <= 0 {
		fmt.Println("Invalid Product ID. Please enter a positive integer.")
		return
	}

	fmt.Print("Enter Product Name: ")
	scanner.Scan()
	name := scanner.Text()
	if strings.TrimSpace(name) == "" {
		fmt.Println("Product Name cannot be empty.")
		return
	}

	fmt.Print("Enter Product Price: ")
	scanner.Scan()
	price, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil || price < 0 {
		fmt.Println("Invalid Price. Please enter a non-negative number.")
		return
	}

	fmt.Print("Enter Product Stock: ")
	scanner.Scan()
	stock, err := strconv.Atoi(scanner.Text())
	if err != nil || stock < 0 {
		fmt.Println("Invalid Stock. Please enter a non-negative integer.")
		return
	}

	err = services.AddProduct(id, name, price, stock)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Product added successfully.")
	}
}

func UpdateStockHandler(scanner *bufio.Scanner) {
	fmt.Print("Enter Product ID: ")
	scanner.Scan()
	id, err := strconv.Atoi(scanner.Text())
	if err != nil || id <= 0 {
		fmt.Println("Invalid Product ID. Please enter a positive integer.")
		return
	}

	fmt.Print("Enter New Stock: ")
	scanner.Scan()
	newStock, err := strconv.Atoi(scanner.Text())
	if err != nil || newStock < 0 {
		fmt.Println("Invalid Stock. Please enter a non-negative integer.")
		return
	}

	err = services.UpdateStock(id, newStock)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Stock updated successfully.")
	}
}

func SearchProductHandler(scanner *bufio.Scanner) {
	fmt.Print("Enter Product Name or ID to search: ")
	scanner.Scan()
	query := scanner.Text()

	product, err := services.SearchProduct(query)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Product Found: %+v\n", product)
	}
}

func DisplayInventoryHandler() {
	inventory := services.GetInventory()
	if len(inventory) == 0 {
		fmt.Println("Inventory is empty.")
		return
	}

	fmt.Println("ID\tName\t\tPrice\t\tStock")
	fmt.Println("--------------------------------------------------")
	for _, product := range inventory {
		fmt.Printf("%d\t%s\t\t%.2f\t\t%d\n", product.ID, product.Name, product.Price, product.Stock)
	}
}

func SortInventoryHandler(scanner *bufio.Scanner) {
	fmt.Print("Enter sorting criteria ('price' or 'stock'): ")
	scanner.Scan()
	criteria := scanner.Text()

	sortedInventory, err := services.SortInventory(criteria)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("ID\tName\t\tPrice\t\tStock")
	fmt.Println("--------------------------------------------------")
	for _, product := range sortedInventory {
		fmt.Printf("%d\t%s\t\t%.2f\t\t%d\n", product.ID, product.Name, product.Price, product.Stock)
	}
}
