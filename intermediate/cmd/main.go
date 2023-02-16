package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"users/pkg/storage"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Get("/show/{city_id}", storage.GetRecordById)
	r.Post("/create", storage.Create)
	r.Delete("/delete", storage.Delete)
	r.Put("/update/{city_id}", storage.Update)

	r.Route("/list-cities", func(r chi.Router) {
		r.Get("/", storage.GetAll)
		r.Post("/by-option", storage.ListCities)
	})

	go func() {
		http.ListenAndServe("localhost:8080", r)
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// <-sigChan
	// fmt.Println("Done!", <-sigChan)

	for {
		select {
		case <-sigChan:
			storage.ShutDown()
			return
		}
	}

}
