package storage

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"users/pkg/model"

	"github.com/go-chi/chi"
	"github.com/gocarina/gocsv"
)

var storage []*model.City

func init() {
	// Читаем файл CSV с информацией о городах
	file, err := os.OpenFile("data/cities.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if err := gocsv.UnmarshalFile(file, &storage); err != nil {
		panic(err)
	}
}

// GetRecordById - Получение полной информации о городе по его id
func GetRecordById(w http.ResponseWriter, r *http.Request) {
	cityId, _ := strconv.Atoi(fmt.Sprintf("%v", chi.URLParam(r, "city_id")))
	id := uint64(cityId)

	name := "NotFound"
	for _, city := range storage {
		if uint64(city.Id) == id {
			name = city.Name
			break
		}
	}

	// w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Нашли город:\n" + string(name) + "\n"))
	return
}

// GetAll - Возвращает полную информаци по всем городам.
func GetAll(w http.ResponseWriter, r *http.Request) {
	var response string

	for _, city := range storage {
		response += toString(city)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
	return
}

// ***
// Хелперы

func toString(city *model.City) string {
	return strconv.Itoa(int(city.Id)) + "\t" + city.Name + "\n"
}
