package main

import (
	"flag"
	"net/http"
	"strconv"
	"strings"

	"skillbox/pkg/logs"
	user "skillbox/pkg/user"

	"github.com/go-chi/chi"
	"github.com/spf13/viper"
)

func init() {
	logs.InitialSet("main")
}

func main() {
	v := viper.New()
	v.SetConfigFile("config/config.yml")
	v.ReadInConfig()

	server := ""
	flag.StringVar(&server, "s", "first", "which server up of two")
	flag.Parse()

	host := v.GetString("server." + server + ".host")
	port := v.GetInt("server." + server + ".port")

	portStr := strconv.Itoa(port)
	lis := host + ":" + portStr

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

	http.ListenAndServe(lis, r)
}

func Mid(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		server := r.Host
		serverSplit := strings.Split(server, ":")
		port := serverSplit[1]

		if port == "8081" {
			server = "second"
		} else {
			server = "first"
		}

		w.Write([]byte("-> the " + server + " server speaks\n"))
		next.ServeHTTP(w, r)
	})
}
