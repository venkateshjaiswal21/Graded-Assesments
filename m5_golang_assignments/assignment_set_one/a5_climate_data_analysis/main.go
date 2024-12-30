package main

import (
	"a5_climate_data_analysis/handlers"
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nClimate Data Analysis System")
		fmt.Println("1. Add New City")
		fmt.Println("2. Show Highest Temperature City")
		fmt.Println("3. Show Lowest Temperature City")
		fmt.Println("4. Show Average Rainfall")
		fmt.Println("5. Filter Cities by Rainfall")
		fmt.Println("6. Show All Cities")
		fmt.Println("7. Exit")
		fmt.Print("\nEnter your choice (1-7): ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			handlers.AddCityHandler(scanner)
		case "2":
			handlers.DisplayHighestTemperatureHandler()
		case "3":
			handlers.DisplayLowestTemperatureHandler()
		case "4":
			handlers.DisplayAverageRainfallHandler()
		case "5":
			handlers.FilterCitiesByRainfallHandler(scanner)
		case "6":
			handlers.DisplayAllCitiesHandler()
		case "7":
			fmt.Println("\nExiting program. Goodbye!")
			return
		default:
			fmt.Println("\nInvalid choice. Please enter a number between 1 and 7.")
		}
	}
}
