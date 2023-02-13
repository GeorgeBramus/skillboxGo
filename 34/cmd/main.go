package main

import (
	"users/pkg/storage"
)

func main() {
	// user1 := user.User{
	// 	Name: "Вася",
	// 	Age:  22,
	// }

	// user2 := user.User{
	// 	Name: "Миша",
	// 	Age:  23,
	// }

	// users := []*user.User{}
	// users = append(users, &user1, &user2)

	storage.Get()
}
