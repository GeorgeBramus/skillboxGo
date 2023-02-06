package main

import (
	"net/http"

	"skillbox/pkg/logs"
	user "skillbox/pkg/user"

	"github.com/go-chi/chi"
	// log "github.com/sirupsen/logrus"
)

func init() {
	logs.InitialSet("main")
}

func main() {
	// db := database.Initial()
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Post("/create", user.Create)
	r.Post("/make_friends", user.MakeFriends)
	r.Delete("/user", user.DeleteUser)
	r.Get("/friends/{user_id}", user.Friends)

	r.Get("/get", user.GetAll)

	// name := "George"
	// if id, err := user.Find(name); err != nil {
	// 	log.Error(err)
	// 	fmt.Println("Пользователь", name, "не найден")
	// } else {
	// 	fmt.Println("ID пользователя", name+":", id)
	// }

	http.ListenAndServe("localhost:8080", r)

}
