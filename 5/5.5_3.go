package main

import "fmt"

func main() {
	var (
		firstNum, secondNum, thirdNum int
		inputError                    int = 1
		isCheck                       bool
	)

	for {
		switch inputError {
		case 1:
			fmt.Println("Введите три разных числа через пробел:")
		case 2:
			fmt.Println("Ещё одна попытка! Числа должны быть разными! (вводить через пробел)")
		case 3:
			fmt.Println("Это последняя попытка(вводить через пробел):")
		}
		fmt.Scan(&firstNum, &secondNum, &thirdNum)

		if firstNum != secondNum && firstNum != thirdNum && secondNum != thirdNum {
			isCheck = true
			break
		} else if inputError >= 3 {
			break
		} else {
			inputError++
		}
	}

	if isCheck {
		fmt.Println("Все числа разные!")
	} else {
		fmt.Println("Не удалось с трёх попыток ввести разные числа.")
	}
}
