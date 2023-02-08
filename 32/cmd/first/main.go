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

	r.Use(Mid)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Post("/create", user.Create)
	r.Post("/make_friends", user.MakeFriends)
	r.Delete("/user", user.DeleteUser)
	r.Get("/friends/{user_id}", user.Friends)
	r.Put("/{user_id}", user.Update)

	r.Get("/get", user.GetAll)

	http.ListenAndServe("localhost:8080", r)
}

func Mid(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("-> Ответил 1й сервер\n"))
		next.ServeHTTP(w, r)
	})
}
