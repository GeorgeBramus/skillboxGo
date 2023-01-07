package main

import "fmt"

func isEvenNumber(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}

func main() {
	var (
		num     int
		nameNum string = "нечётное"
	)
	fmt.Print("Введите число: ")
	fmt.Scan(&num)
	if isEvenNumber(num) {
		nameNum = "чётное"
	}
	fmt.Printf("Вы ввели %v число.\n", nameNum)
}
