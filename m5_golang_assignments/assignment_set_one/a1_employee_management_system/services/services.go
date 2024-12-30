package services

import (
	"a1_employee_management_system/models"
	"errors"
	"fmt"
	"strings"
)

var employees []models.Employee

func AddEmployee(id int, name string, age int, department string) error {
	if id <= 0 {
		return errors.New("ID must be a positive integer")
	}
	if len(name) == 0 || strings.TrimSpace(name) == "" {
		return errors.New("name must be a non-empty string")
	}
	if age <= 18 {
		return errors.New("age must be greater than 18")
	}
	if len(department) == 0 || strings.TrimSpace(department) == "" {
		return errors.New("department must be a non-empty string")
	}
	for _, emp := range employees {
		if emp.ID == id {
			return errors.New("employee ID must be unique")
		}
	}

	employees = append(employees, models.Employee{
		ID:         id,
		Name:       name,
		Age:        age,
		Department: department,
	})
	return nil
}

func SearchEmployee(query string) (*models.Employee, error) {
	for _, emp := range employees {
		if fmt.Sprintf("%d", emp.ID) == query || strings.EqualFold(emp.Name, query) {
			return &emp, nil
		}
	}
	return nil, errors.New("employee not found")
}

func ListEmployeesByDepartment(department string) []models.Employee {
	var result []models.Employee
	for _, emp := range employees {
		if strings.EqualFold(emp.Department, department) {
			result = append(result, emp)
		}
	}
	return result
}

func CountEmployeesByDepartment(department string) int {
	count := 0
	for _, emp := range employees {
		if strings.EqualFold(emp.Department, department) {
			count++
		}
	}
	return count
}
