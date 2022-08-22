package main

import "fmt"

func main() {
	var (
		firstNum, secondNum, thirdNum int
		notification                  string = "нет числа"
	)

	fmt.Println("Три числа.")

	whatNumberName := [3]string{"первое", "второе", "третье"}
	for i := 0; i < 3; i++ {
		fmt.Println("Введите", whatNumberName[i], "число")
		switch i {
		case 0:
			fmt.Scan(&firstNum)
		case 1:
			fmt.Scan(&secondNum)
		case 2:
			fmt.Scan(&thirdNum)
		}
	}
	if firstNum > 5 || secondNum > 5 || thirdNum > 5 {
		notification = "есть число"
	}

	fmt.Println("Среди введённых чисел", notification, "больше 5.")
}

/*
	Link to Repl IT
	https://replit.com/@GieorghiiZaporo/Three-numbers?v=1
*/
