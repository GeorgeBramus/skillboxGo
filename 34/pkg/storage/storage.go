package storage

import (
	"fmt"
	"users/pkg/user"
)

var storage []*user.User

func init() {
	user1 := user.User{
		Name: "Вася",
		Age:  22,
	}

	user2 := user.User{
		Name: "Миша",
		Age:  23,
	}

	// users := []*user.User{}
	storage = append(storage, &user1, &user2)
}

func Get() {
	for _, user := range storage {
		fmt.Println(user.Name)

	}
}
