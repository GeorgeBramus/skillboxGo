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

// Delete - Удаление информации о городе по id
func Delete(w http.ResponseWriter, r *http.Request) {
	// Читаю запрос
	// {"city_id":1}
	content, err := ioutil.ReadAll((r.Body))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("> " + err.Error()))
		return
	}
	defer r.Body.Close()

	type delete struct {
		Id uint64 `json:"city_id"`
	}

	var cityForRemoval delete

	if err := json.Unmarshal(content, &cityForRemoval); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	var (
		cityFound bool
		cityName  string
		cityId    uint64
	)

	for i, city := range storage {
		if city.Id == cityForRemoval.Id {
			cityFound = true
			cityName = city.Name
			cityId = city.Id
			storage = removalFromSlice(storage, i)
			break
		}
	}

	if !cityFound {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Город с таким ID не найден\n"))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Информация о городе " + cityName + " с ID[" + strconv.Itoa(int(cityId)) + "] была удалена.\n"))
	return
}

// Update - Обновение информации о численности населения города по указанному id
func Update(w http.ResponseWriter, r *http.Request) {
	cityId, _ := strconv.Atoi(fmt.Sprintf("%v", chi.URLParam(r, "city_id")))
	id := uint64(cityId)

	// Читаю запрос
	// {"new_population":1234567}
	content, err := ioutil.ReadAll((r.Body))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	type Update struct {
		Population uint `json:"new_population"`
	}

	var update Update

	if err := json.Unmarshal(content, &update); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	var (
		cityFound                    bool
		oldPopulation, newPopulation uint
	)

	for _, city := range storage {
		if city.Id == id {
			cityFound = true
			oldPopulation = city.Population
			city.Population = update.Population
			newPopulation = city.Population
			break
		}
	}

	if !cityFound {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Город с таким ID не найден\n"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Информация Population обновллена\nБыло: " + strconv.Itoa(int(oldPopulation)) + "\nСтало: " + strconv.Itoa(int(newPopulation)) + "\n"))
	return
}

// ListCities - Возвращает список городов по указанному признаку.
// Варианты запросов:
// 1) {"region":"region_name"} города с таким регионои
// 2) {"district":"district_name"} города с таким округом
// 3) {"population":[{"from":1,"to":2}]} диапазон численности
// 4) {"foundation":[{"from":1,"to":2}]} диапазон по году основания
func ListCities(w http.ResponseWriter, r *http.Request) {
	// Читаю запрос
	content, err := ioutil.ReadAll((r.Body))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	type Range struct {
		From uint `json:"from,omitempty"`
		To   uint `json:"to,omitempty"`
	}

	type changeOptions struct {
		Region     string `json:"region,omitempty"`
		District   string `json:"district,omitempty"`
		Population Range  `json:"population,omitempty"`
		Foundation Range  `json:"foundation,omitempty"`
	}

	var options changeOptions

	if err := json.Unmarshal(content, &options); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	var maxPopulation uint
	for _, city := range storage {
		if city.Population > maxPopulation {
			maxPopulation = city.Population
		}
	}
	var maxFoundation uint
	for _, city := range storage {
		if city.Foundation > maxFoundation {
			maxFoundation = city.Foundation
		}
	}

	response := ""
	// неправильный запрос, или правильный, но ничего нешли?
	var correctRequest bool
	// Предполагаю, что будет использован один из вариантов запроса.
	// Поэтому определяем запрос по первому не пустому полю.
	if options.Region != "" {
		for _, city := range storage {
			if city.Region == options.Region {
				response += toStr(city)
			}
		}
		correctRequest = true
	} else if options.District != "" {
		for _, city := range storage {
			if city.District == options.District {
				response += toStr(city)
			}
		}
		correctRequest = true
	} else if options.Population.From > 0 || options.Population.To > 0 {
		if options.Population.To < options.Population.From {
			options.Population.To = maxPopulation
		}
		for _, city := range storage {
			if city.Population >= options.Population.From && city.Population <= options.Population.To {
				response += toStr(city)
			}
		}
		correctRequest = true
	} else if options.Foundation.From > 0 || options.Foundation.To > 0 {
		if options.Foundation.To < options.Foundation.From {
			options.Foundation.To = maxFoundation
		}
		for _, city := range storage {
			if city.Foundation >= options.Foundation.From && city.Foundation <= options.Foundation.To {
				response += toStr(city)
			}
		}
		correctRequest = true
	}
	if correctRequest {
		response = "По ваашему запросу ничего не найдено\n"
	} else {
		response = "Некорректый запрос\n"
	}

	w.Header().Add("Content-Type", "application/json")
	if correctRequest {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Write([]byte(response))
	return
}

// ShutDown - Завершение программы и запись из оперативной памяти в файл
// изменённой и новой информации по городам
func ShutDown() {
	cities := storage

	// Читаем файл CSV с информацией о городах
	file, err := os.OpenFile("data/cities.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	gocsv.MarshalFile(&cities, file)
}

// ***
// Хелперы

func toStr(city *model.City) string {
	return strconv.Itoa(int(city.Id)) + "\t" + city.Name + "\t" + city.Region + "\t" + city.District + "\t" + strconv.Itoa(int(city.Population)) + "\t" + strconv.Itoa(int(city.Foundation)) + "\n"
}

func toRow(city *model.City) string {
	return "<tr><td>" + strconv.Itoa(int(city.Id)) + "</td><td>" + city.Name + "</td><td>" + city.Region + "</td><td>" + city.District + "</td><td>" + strconv.Itoa(int(city.Population)) + "</td><td>" + strconv.Itoa(int(city.Foundation)) + "</td></tr>"
}

func removalFromSlice(slice []*model.City, index int) []*model.City {
	slice[len(slice)-1], slice[index] = slice[index], slice[len(slice)-1]
	return slice[:len(slice)-1]
}
