package storage

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"city/pkg/city"

	"github.com/go-chi/chi"
	"github.com/gocarina/gocsv"
)

type Storage []*city.City

// var storage Storage
// storage, _ = New()

// New - Новое хранилище с начальными данными
func New() (Storage, error) {
	// Читаем файл CSV с информацией о городах
	citiesFile, err := os.OpenFile("data/cities.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer citiesFile.Close()

	cities := Storage{}
	if err := gocsv.UnmarshalFile(citiesFile, &cities); err != nil {
		return nil, err
	}
	return cities, nil
}

// GetInfo - Получение информации о городе по его id
func GetInfo(w http.ResponseWriter, r *http.Request) {
	cityId, _ := strconv.Atoi(fmt.Sprintf("%v", chi.URLParam(r, "id")))
	id := uint64(cityId)

	storage := &Storage{}

	for _, city := range storage {
		if uint64(city.Id) == id {
			name := city.Name
			break
		}
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Нашли город:\n" + name + "\n"))
	return
}

// func (storage Storage) Put(student *student.Students) {
// 	storage[student.Name] = student
// }

// func (storage Storage) Get(studentName string) (*student.Students, error) {
// 	stud, ok := storage[studentName]
// 	if !ok {
// 		return nil, fmt.Errorf("no such user")
// 	}
// 	return stud, nil
// }
