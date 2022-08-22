package main

import "fmt"

func main() {
	var (
		numberTicket   int
		isCorrectInput bool
	)

	fmt.Println("Укажите четырёхзначный номер Вашего билета:")
	for !isCorrectInput {
		fmt.Scan(&numberTicket)
		if numberTicket > 999 && numberTicket <= 9999 {
			isCorrectInput = true
		} else {
			fmt.Println("Номер должен быть четырёхзначным! Попробуйте ещё раз:")
		}
	}

	divider := 10
	remainder := []int{}
	for i := 0; i < 4; i++ {
		remainder = append(remainder, numberTicket%divider)
		numberTicket /= divider
	}

	// Если зеркальный билет
	if remainder[0] == remainder[3] && remainder[1] == remainder[2] {
		fmt.Println("Билет является зеркальным.")
	} else if remainder[0]+remainder[1] == remainder[2]+remainder[3] {
		fmt.Println("У Вас счастливый билет.")
	} else {
		fmt.Println("У Вас обычный билет.")
	}
}
