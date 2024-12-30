package services

import (
	"a5_climate_data_analysis/models"
	"strings"
)

var cities []models.City

func AddCity(city models.City) {
	cities = append(cities, city)
}

func FindHighestTemperature() (models.City, bool) {
	if len(cities) == 0 {
		return models.City{}, false
	}
	highest := cities[0]
	for _, city := range cities[1:] {
		if city.Temperature > highest.Temperature {
			highest = city
		}
	}
	return highest, true
}

func FindLowestTemperature() (models.City, bool) {
	if len(cities) == 0 {
		return models.City{}, false
	}
	lowest := cities[0]
	for _, city := range cities[1:] {
		if city.Temperature < lowest.Temperature {
			lowest = city
		}
	}
	return lowest, true
}

func CalculateAverageRainfall() (float64, bool) {
	if len(cities) == 0 {
		return 0, false
	}

	total := 0.0
	for _, city := range cities {
		total += city.Rainfall
	}
	return total / float64(len(cities)), true
}

func FilterCitiesByRainfall(threshold float64) []models.City {
	var filteredCities []models.City
	for _, city := range cities {
		if city.Rainfall > threshold {
			filteredCities = append(filteredCities, city)
		}
	}
	return filteredCities
}

func SearchCity(name string) (models.City, bool) {
	searchName := strings.ToLower(name)
	for _, city := range cities {
		if strings.ToLower(city.Name) == searchName {
			return city, true
		}
	}
	return models.City{}, false
}

func GetAllCities() []models.City {
	return cities
}
