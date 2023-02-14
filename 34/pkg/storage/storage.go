package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

// GetAll - Возвращает полную информаци по всем городам.
func GetAll(w http.ResponseWriter, r *http.Request) {
	var response string = "<table>"
	response += "<tr><th>Id</th><th>Name</th><th>Region</th><th>District</th><th>Population</th><th>Foundation</th></tr>"

	for _, city := range storage {
		response += toRow(city)
	}
	response += "</table"

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
	return
}

// GetRecordById - Получение полной информации о городе по его id
func GetRecordById(w http.ResponseWriter, r *http.Request) {
	cityId, _ := strconv.Atoi(fmt.Sprintf("%v", chi.URLParam(r, "city_id")))
	id := uint64(cityId)

	var cityFound *model.City
	cityNotFound := true

	for _, city := range storage {
		if uint64(city.Id) == id {
			cityFound = city
			cityNotFound = false
			break
		}
	}

	if cityNotFound {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Город с такимм ID не найден."))
		return
	}

	// w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(int(cityFound.Id)) + "\t" + cityFound.Name + "\t" + cityFound.Region + "\t" + cityFound.District + "\t" + strconv.Itoa(int(cityFound.Population)) + "\t" + strconv.Itoa(int(cityFound.Foundation)) + "\n"))
	return
}

// Create - Создание новой записи в хранилище
func Create(w http.ResponseWriter, r *http.Request) {
	// Читаю запрос
	// {"name":"cityName",
	//  "region":"regionName",
	//  "district":"districtName",
	//  "population":123456,
	//  "foundation":1234}
	// id проставляем самостоятельно
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	// Разбираю JSON в структуру
	var city *model.City
	if err := json.Unmarshal(content, &city); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Ищем максимальный id в хранилище
	city.Id = 0
	for _, cityFromStorage := range storage {
		if uint64(cityFromStorage.Id) > city.Id {
			city.Id = uint64(cityFromStorage.Id)
		}
	}
	// ..и прибавляем единицу
	city.Id++
	// ..и полную информацию о новом городе добавим в хранилище
	storage = append(storage, city)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Информация о новом городе добавлена в хранилище\n"))
	return
}

// ***
// Хелперы

func toRow(city *model.City) string {
	return "<tr><td>" + strconv.Itoa(int(city.Id)) + "</td><td>" + city.Name + "</td><td>" + city.Region + "</td><td>" + city.District + "</td><td>" + strconv.Itoa(int(city.Population)) + "</td><td>" + strconv.Itoa(int(city.Foundation)) + "</td></tr>"
}
