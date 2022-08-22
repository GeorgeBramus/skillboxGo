package main

import "fmt"

func main() {
	var firstNum, secondNum, thirdNum int
	fmt.Println("Введите три числа через пробел:")
	fmt.Scan(&firstNum, &secondNum, &thirdNum)

	if firstNum > 0 || secondNum > 0 || thirdNum > 0 {
		fmt.Println("Одно из чисел является положительным.")
	} else {
		fmt.Println("Ни одно из чисел не является положительным.")
	}
}
