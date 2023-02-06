package actions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"skillbox/pkg/database"
	"skillbox/pkg/logs"
	"skillbox/pkg/model"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

type user struct {
	Name    string   `json:"name"`
	Age     uint8    `json:"age"`
	Friends []string `json:"friends"`
}

func init() {
	logs.InitialSet("main")
}

// Create Создание юзера
func Create(w http.ResponseWriter, r *http.Request) {
	db := database.Initial()
	// Читаю запрос
	// {"name":"some name","age":24,"friends":[]}
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	// Разбираю JSON в структуру
	var user user
	if err := json.Unmarshal(content, &user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Возьму из JSON некоторые поля для записи юзера в БД
	uRecord := model.User{
		Name: user.Name,
		Age:  user.Age,
	}

	// Собственно теперь запишу эти поля в БД
	result := db.Create(&uRecord)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(result.Error.Error()))
		return
	}

	// Если у пользователя есть друзья,
	// указанные в JSON, то и их надо записать
	if len(user.Friends) > 0 {
		friendsId := friendsIdArray(user)
		// Подготовим структуру для записи в базу
		var userAndHisFriends []model.Friends
		userId := uint64(uRecord.Id)
		for _, friendId := range friendsId {
			comWithAFriend := model.Friends{
				UserId:   userId,
				FriendId: friendId,
			}
			userAndHisFriends = append(userAndHisFriends, comWithAFriend)
		}

		// Теперь можно записать в базу друзей пользователя
		db.Create(&userAndHisFriends)
	}

	// Теперь можно отпавить ответ
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Пользователь " + user.Name + " создан с ID[" + strconv.Itoa(int(uRecord.Id)) + "]\n"))

	return
}

// MakeFriends - Подружить двух пользователей
func MakeFriends(w http.ResponseWriter, r *http.Request) {
	db := database.Initial()
	// Читаю запрос
	// {"source_id":1,"target_id":2}
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	type friends struct {
		UserId   uint64 `json:"source_id"`
		FriendId uint64 `json:"target_id"`
	}

	var f friends

	if err := json.Unmarshal(content, &f); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	result := db.Create(f)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	} else {
		f.UserId, f.FriendId = f.FriendId, f.UserId
		db.Create(f)
	}

	userName, _ := findName(f.FriendId)
	friendName, _ := findName(f.UserId)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Пользователи " + userName + " и " + friendName + " стали друзьями\n"))
	return
}

// DeleteUser Удаление юзера по id
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db := database.Initial()
	// Читаю запрос
	// {"target_id":1}
	content, err := ioutil.ReadAll((r.Body))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("> " + err.Error()))
		return
	}
	defer r.Body.Close()

	type dUser struct {
		TargetId uint64 `json:"target_id"`
	}

	var user dUser

	if err := json.Unmarshal(content, &user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	userName, userExistsErr := findName(uint64(user.TargetId))
	if userExistsErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Пользователь не найден\n"))
		return
	}

	var (
		u              model.User
		f              model.Friends
		friendsDeleted bool
	)

	if result := db.Where("id = ?", user.TargetId).Delete(&u); result.Error != nil {
		log.Error("Не удалось удалить запись в таблице пользователей. Пользователь: " + userName + ". ID пользователя: " + strconv.Itoa(int(user.TargetId)))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Пользователь не был удалён.\n"))
		return
	} else {
		// Удаляем всех друзей у пользователя, которого только что удалили
		if result := db.Where("user_id = ?", user.TargetId).Delete(&f); result.Error != nil {
			log.Error("Не удалось удалить записи друзей пользователя, которого удалили. Пользователь: " + userName + ". ID пользователя: " + strconv.Itoa(int(user.TargetId)))
			friendsDeleted = true
		}
		// И теперь можно удалить записи из таблицы Friends,
		// где этот пользователь был чьим-то другом.
		if result := db.Where("friend_id = ?", user.TargetId).Delete(&f); result.Error != nil {
			log.Error("Не удалось удалить записи друзей пользователя, которого удалили. Пользователь: " + userName + ". ID пользователя: " + strconv.Itoa(int(user.TargetId)))
			friendsDeleted = true
		}

	}

	messageAboutRemovingFriends := ""
	if friendsDeleted {
		messageAboutRemovingFriends = "[!] Не были удалены записи с друзьями пользователя и/или записи, где пользователь состоит в друзьях других пользователей.\n"
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Пользователь " + userName + " удалён\n" + messageAboutRemovingFriends))

	return
}

// Friends - Показать всех друзей конкретного пользователя
func Friends(w http.ResponseWriter, r *http.Request) {
	db := database.Initial()
	// Читаю запрос
	userID, _ := strconv.Atoi(fmt.Sprintf("%v", chi.URLParam(r, "user_id")))
	id := uint64(userID)

	userName, userExistsErr := findName(id)
	if userExistsErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Пользователь не найден\n"))
		return
	}

	// type Friends struct {
	// 	FriendId uint64 `json:"friend_id"`
	// }
	var f []model.Friends

	db.Where("user_id = ?", id).Find(&f)

	userFriends := ""
	for _, friend := range f {
		friendName, _ := findName(friend.FriendId)
		userFriends += " - " + friendName + "\n"
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Друзья пользователя " + userName + "\n" + userFriends + "\n"))
	return
}

// GetAll - Получение всех пользователей
func GetAll(w http.ResponseWriter, r *http.Request) {
	db := database.Initial()
	var users []model.User
	result := db.Find(&users)
	var response string

	if result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(response))
		return
	}

	for _, user := range users {
		response += toString(user)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
	return
}

/***
 * Вспомогательные функции
 */

// findName - Поиск по id и возвращает имени
func findName(id uint64) (string, error) {
	db := database.Initial()
	var u model.User
	if user := db.First(&u, id); user.Error != nil {
		return "", user.Error
	} else {
		return u.Name, nil
	}
}

// findId - Поиск по имени и возврат id найденного юзера
func findId(name string) (uint64, error) {
	db := database.Initial()
	var u model.User
	if user := db.Where("name = ?", name).First(&u); user.Error != nil {
		return 0, user.Error
	} else {
		return uint64(u.Id), nil
	}
}

// friendsId - Соберём массив id друзей, если они есть в базе.
// Если таких нет в базе, то не будут добавлены к друзьям юзера. Такой вот косяк.
func friendsIdArray(u user) []uint64 {
	var friendsId []uint64
	for _, friend := range u.Friends {
		id, err := findId(friend)
		if err != nil {
			// Скорее всего друга просто нет в базе
			log.WithFields(log.Fields{
				"user_name":   u.Name,
				"friend_name": friend,
			}).Warn("При попытке добавления друга пользоователя произошла ошибка")
			continue
		}
		friendsId = append(friendsId, id)
	}
	return friendsId
}

// toString - Информация о пользователе в виде строки
func toString(u model.User) string {
	return fmt.Sprintf("ID: %d\nИмя: %s\nВозраст: %d\n\n", u.Id, u.Name, u.Age)
}
