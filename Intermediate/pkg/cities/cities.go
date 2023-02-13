package cities

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/gocarina/gocsv"
)

type Cities struct {
	Id         uint64 `csv:"id"`
	Name       string `csv:"name"`
	Region     string `csv:"region"`
	District   string `csv:"district"`
	Population uint   `csv:"population"`
	Foundation uint16 `csv:"foundation"`
}

// Инициализация
func New(citiesFile *os.File) ([]*Cities, error) {
	cities := []*Cities{}
	if err := gocsv.UnmarshalFile(citiesFile, &cities); err != nil {
		return nil, err
	}
	return cities, nil
}

func GetInfo(w http.ResponseWriter, r *http.Request) {
	// ожидается id города
	// {"id": 744}
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	cityId, _ := strconv.Atoi(fmt.Sprintf("%v", chi.URLParam(r, "id")))
	id := uint64(cityId)

	// for _, city := range cities {
	// 	if uint64(city.Id) == id {
	// 		return city, nil
	// 	}
	// }
	// return nil, fmt.Errorf("Город с идентификатором %v не найден.", id)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Пользователи " + userName + " и " + friendName + " стали друзьями\n"))
	return
}

// func main() {

// for _, city := range cities {
// 	fmt.Println("Название города:", city.Name)
// }

// if _, err := clientsFile.Seek(0, 0); err != nil { // Go to the start of the file
// 	panic(err)
// }

// clients = append(clients, &Client{Id: "12", Name: "John", Age: "21"}) // Add clients
// clients = append(clients, &Client{Id: "13", Name: "Fred"})
// clients = append(clients, &Client{Id: "14", Name: "James", Age: "32"})
// clients = append(clients, &Client{Id: "15", Name: "Danny"})
// csvContent, err := gocsv.MarshalString(&clients) // Get all clients as CSV string
// //err = gocsv.MarshalFile(&clients, clientsFile) // Use this to save the CSV back to the file
// if err != nil {
// 	panic(err)
// }
// fmt.Println(csvContent) // Display all clients as CSV string
// }
