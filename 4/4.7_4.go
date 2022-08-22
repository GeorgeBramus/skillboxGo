package main

import "fmt"

func main() {
	var firstNum, secondNum, thirdNum int

	fmt.Println("Три числа.")

	whatNumberName := [3]string{"первое", "второе", "третье"}
	count := 0
	for i := 0; i < 3; i++ {
		fmt.Println("Введите", whatNumberName[i], "число")
		switch i {
		case 0:
			fmt.Scan(&firstNum)
			if firstNum >= 5 {
				count++
			}
		case 1:
			fmt.Scan(&secondNum)
			if secondNum >= 5 {
				count++
			}
		case 2:
			fmt.Scan(&thirdNum)
			if thirdNum >= 5 {
				count++
			}
		}
	}

	fmt.Println("Среди введённых чисел", count, "больше или равны 5.")
}

/*
	Link to Repl IT
	https://replit.com/@GieorghiiZaporo/three-numbers-2?v=1
*/
