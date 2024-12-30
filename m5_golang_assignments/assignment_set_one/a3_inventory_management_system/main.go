package main

import (
	"a3_inventory_management_system/handlers"
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nInventory Management System")
		fmt.Println("1. Add Product")
		fmt.Println("2. Update Stock")
		fmt.Println("3. Search Product")
		fmt.Println("4. Display Inventory")
		fmt.Println("5. Sort Inventory")
		fmt.Println("6. Exit")
		fmt.Print("\nEnter your choice (1-6): ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			handlers.AddProductHandler(scanner)
		case "2":
			handlers.UpdateStockHandler(scanner)
		case "3":
			handlers.SearchProductHandler(scanner)
		case "4":
			handlers.DisplayInventoryHandler()
		case "5":
			handlers.SortInventoryHandler(scanner)
		case "6":
			fmt.Println("Exiting... Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
