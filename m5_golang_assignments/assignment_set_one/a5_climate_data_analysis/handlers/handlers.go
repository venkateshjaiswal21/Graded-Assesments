package handlers

import (
	"a5_climate_data_analysis/models"
	"a5_climate_data_analysis/services"
	"bufio"
	"fmt"
	"strconv"
)

func AddCityHandler(scanner *bufio.Scanner) {
	fmt.Print("\nEnter city name: ")
	scanner.Scan()
	name := scanner.Text()

	if name == "" {
		fmt.Println("City name cannot be empty")
		return
	}

	fmt.Print("Enter temperature (°C): ")
	scanner.Scan()
	var temp float64
	_, err := fmt.Sscanf(scanner.Text(), "%f", &temp)
	if err != nil {
		fmt.Println("Invalid temperature. Please enter a valid number.")
		return
	}

	fmt.Print("Enter rainfall (mm): ")
	scanner.Scan()
	var rainfall float64
	_, err = fmt.Sscanf(scanner.Text(), "%f", &rainfall)
	if err != nil {
		fmt.Println("Invalid rainfall. Please enter a valid number.")
		return
	}

	services.AddCity(models.City{Name: name, Temperature: temp, Rainfall: rainfall})
	fmt.Println("City added successfully!")
}

func DisplayHighestTemperatureHandler() {
	if highest, found := services.FindHighestTemperature(); found {
		fmt.Printf("\nCity with highest temperature: %s (%.1f°C)\n", highest.Name, highest.Temperature)
	} else {
		fmt.Println("\nNo cities in database")
	}
}

func DisplayLowestTemperatureHandler() {
	if lowest, found := services.FindLowestTemperature(); found {
		fmt.Printf("\nCity with lowest temperature: %s (%.1f°C)\n", lowest.Name, lowest.Temperature)
	} else {
		fmt.Println("\nNo cities in database")
	}
}

func DisplayAverageRainfallHandler() {
	if avg, found := services.CalculateAverageRainfall(); found {
		fmt.Printf("\nAverage rainfall across all cities: %.1f mm\n", avg)
	} else {
		fmt.Println("\nNo cities in database")
	}
}

func FilterCitiesByRainfallHandler(scanner *bufio.Scanner) {
	fmt.Print("\nEnter rainfall threshold (mm): ")
	scanner.Scan()
	threshold, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}

	filteredCities := services.FilterCitiesByRainfall(threshold)
	if len(filteredCities) == 0 {
		fmt.Printf("\nNo cities found with rainfall above %.1f mm\n", threshold)
	} else {
		fmt.Printf("\nCities with rainfall above %.1f mm:\n", threshold)
		for _, city := range filteredCities {
			fmt.Printf("- %s: %.1f mm\n", city.Name, city.Rainfall)
		}
	}
}

func DisplayAllCitiesHandler() {
	allCities := services.GetAllCities()
	if len(allCities) == 0 {
		fmt.Println("\nNo cities in database")
		return
	}

	fmt.Println("\nAll Cities:")
	for _, city := range allCities {
		fmt.Printf("- %s: %.1f°C, %.1f mm\n", city.Name, city.Temperature, city.Rainfall)
	}
}
