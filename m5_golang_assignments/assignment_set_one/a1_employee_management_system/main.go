package main

import (
	"a1_employee_management_system/handlers"
	"fmt"
)

func main() {
	for {
		fmt.Println("\nEmployee Management System")
		fmt.Println("1. Add Employee")
		fmt.Println("2. Search Employee")
		fmt.Println("3. List Employees by Department")
		fmt.Println("4. Count Employees by Department")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			handlers.AddEmployeeHandler()
		case 2:
			handlers.SearchEmployeeHandler()
		case 3:
			handlers.ListEmployeesByDepartmentHandler()
		case 4:
			handlers.CountEmployeesByDepartmentHandler()
		case 5:
			fmt.Println("Exiting... Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
