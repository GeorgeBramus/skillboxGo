package model

type User struct {
	Name string `csv:"name"`
	Age  int    `csv:"age"`
}
