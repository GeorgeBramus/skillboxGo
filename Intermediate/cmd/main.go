package main

import (
	"city/pkg/storage"
	"fmt"
)

func main() {
	var id uint64 = 744

	// массив структур
	cities, _ := storage.New()

	for _, city := range cities {
		if city.Id == id {
			fmt.Println(city.Name)
			break
		}
	}

	// r := chi.NewRouter()

	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("welcome"))
	// })
	// r.Get("/cities/{city_id}", cities.GetInfo)
}
