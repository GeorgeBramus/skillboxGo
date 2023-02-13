package main

import (
	city "city/pkg/csv"
	"fmt"
	"os"
)

func main() {
	var id uint64 = 744

	citiesFile, err := os.OpenFile("data/cities.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer citiesFile.Close()

	cities, _ := city.NewCities(citiesFile)
	city, _ := city.GetInfo(cities, id)

	fmt.Println(city)
}
