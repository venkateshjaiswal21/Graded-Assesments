package handlers

import (
	"a1_employee_management_system/services"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AddEmployeeHandler() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter ID: ")
	var id int
	fmt.Scanln(&id)

	fmt.Print("Enter Name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter Age: ")
	var age int
	fmt.Scanln(&age)

	fmt.Print("Enter Department: ")
	department, _ := reader.ReadString('\n')
	department = strings.TrimSpace(department)

	if err := services.AddEmployee(id, name, age, department); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Employee added successfully.")
	}
}

func SearchEmployeeHandler() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter ID or Name to search: ")
	query, _ := reader.ReadString('\n')
	query = strings.TrimSpace(query)

	emp, err := services.SearchEmployee(query)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Employee Found: %+v\n", *emp)
	}
}

func ListEmployeesByDepartmentHandler() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Department: ")
	department, _ := reader.ReadString('\n')
	department = strings.TrimSpace(department)

	employees := services.ListEmployeesByDepartment(department)
	if len(employees) == 0 {
		fmt.Println("No employees found in this department.")
	} else {
		fmt.Println("Employees in Department:")
		for _, emp := range employees {
			fmt.Printf("%+v\n", emp)
		}
	}
}

func CountEmployeesByDepartmentHandler() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Department: ")
	department, _ := reader.ReadString('\n')
	department = strings.TrimSpace(department)

	count := services.CountEmployeesByDepartment(department)
	fmt.Printf("Number of employees in %s: %d\n", department, count)
}
