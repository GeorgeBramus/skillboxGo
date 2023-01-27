package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

var ID int64 = 0

type User struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Friends []string `json:"friends"`
}

type Friends struct {
	SourceId int64 `json:"source_id"`
	TargetId int64 `json:"target_id"`
}

func (u *User) toString(id int64) string {
	return fmt.Sprintf("ID: %d\nИмя: %s\nВозраст: %d\nДрузья: %v\n\n", id, u.Name, u.Age, u.Friends)
}

type service struct {
	store map[int64]*User
}

// Create Создание юзера
func (s *service) Create(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadAll((r.Body))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	var u User
	if err := json.Unmarshal(content, &u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	s.store[ID] = &u

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Пользователь " + u.Name + " создан с ID [" + strconv.Itoa(int(ID)) + "]\n"))

	ID++
	return
}

// MakeFriends Подружить двух пользователей
func (s *service) MakeFriends(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadAll((r.Body))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("> " + err.Error()))
		return
	}
	defer r.Body.Close()

	var (
		f            Friends
		sourceFriend *User
		targetFriend *User
	)

	if err := json.Unmarshal(content, &f); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if _, ok := s.store[f.SourceId]; !ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Пользователь не найден\n"))
		return
	}
	if _, ok := s.store[f.TargetId]; !ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Пользователь не найден\n"))
		return
	}

	sourceFriend = s.store[f.SourceId]
	targetFriend = s.store[f.TargetId]

	sourceFriend.Friends = append(sourceFriend.Friends, targetFriend.Name)
	targetFriend.Friends = append(targetFriend.Friends, sourceFriend.Name)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Пользователь " + sourceFriend.Name + " и пользователль " + targetFriend.Name + " стали друзьями\n"))

	return
}

// DeleteUser Удаление юзера по id
func (s *service) DeleteUser(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadAll((r.Body))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("> " + err.Error()))
		return
	}
	defer r.Body.Close()

	type dUser struct {
		TargetId int64 `json:"target_id"`
	}

	var user dUser

	if err := json.Unmarshal(content, &user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if _, ok := s.store[user.TargetId]; !ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Пользователь не найден\n"))
		return
	}

	userForRemoval := s.store[user.TargetId]
	delete(s.store, user.TargetId)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Пользователь " + userForRemoval.Name + " удалён\n"))

	return
}

// GetAll Получение всех пользователей
func (s *service) GetAll(w http.ResponseWriter, r *http.Request) {
	response := ""
	for id, user := range s.store {
		response += user.toString(id)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))

	return
}

// Friends Показать всех друзей конкретного пользователя
func (s *service) Friends(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(fmt.Sprintf("%v", chi.URLParam(r, "user_id")))
	id := int64(userID)

	if _, ok := s.store[id]; !ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Пользователь не найден\n"))
		return
	}

	var user *User
	user = s.store[id]
	userFriends := ""
	for _, friend := range user.Friends {
		userFriends += "&emsp;&bull; " + friend + "\n"
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Друзья пользователя " + user.Name + "\n" + userFriends + "\n"))

	return
}

// Update Обновение данных пользователя
func (s *service) Update(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(fmt.Sprintf("%v", chi.URLParam(r, "user_id")))
	id := int64(userID)

	content, err := ioutil.ReadAll((r.Body))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	type Update struct {
		Age int `json:"new age"`
	}

	var update Update

	if err := json.Unmarshal(content, &update); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if _, ok := s.store[id]; !ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Пользователь не найден\n"))
		return
	}

	var user *User
	user = s.store[id]
	user.Age = update.Age

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Возраст пользователя успешно обновлён\n"))

	return
}

func main() {
	r := chi.NewRouter()
	srv := service{make(map[int64]*User)}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Post("/create", srv.Create)
	r.Post("/make_friends", srv.MakeFriends)
	r.Delete("/user", srv.DeleteUser)
	r.Get("/friends/{user_id}", srv.Friends)
	r.Put("/{user_id}", srv.Update)
	r.Get("/get", srv.GetAll)

	http.ListenAndServe("localhost:8080", r)
}
