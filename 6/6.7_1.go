package main

import "fmt"

func main() {
	var userNum int
	fmt.Print("Введите произвольное целое число: ")
	fmt.Scan(&userNum)

	i := 0
	for i < userNum {
		fmt.Println(i)
		i++
	}
}
