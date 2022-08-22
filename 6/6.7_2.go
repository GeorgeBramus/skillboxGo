package main

import "fmt"

func main() {
	var firstNum, secondNum int
	fmt.Println("Введите два числа для их дальнейшего складывания.\n Второе число должно быть больше первого.")
	fmt.Print("Первое число: ")
	fmt.Scan(&firstNum)
	fmt.Print("Второе число: ")
	fmt.Scan(&secondNum)

	sum := firstNum
	for sum < firstNum+secondNum {
		sum++
	}

	fmt.Println("Результат:", sum)
}
