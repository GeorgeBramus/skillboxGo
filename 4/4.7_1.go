package main

import "fmt"

func main() {
	var firstExamScores, secondExamScores, thirdExamScores, studentScores, finalScores int
	finalScores = 275
	numberName := [3]string{"первого", "второго", "третьего"}

	fmt.Println("Баллы ЕГЭ.")

	for i := 0; i < 3; i++ {
		fmt.Println("Введите результат", numberName[i], "экзамена")
		switch i {
		case 0:
			fmt.Scan(&firstExamScores)
		case 1:
			fmt.Scan(&secondExamScores)
		case 2:
			fmt.Scan(&thirdExamScores)
		}
	}

	fmt.Println("Сумма проходного балла:", finalScores)
	studentScores = firstExamScores + secondExamScores + thirdExamScores
	fmt.Println("Количество набранных быллов:", studentScores)

	var notification string
	if studentScores < finalScores {
		notification = "не "
	}
	fmt.Printf("Вы %sпоступили.\n", notification)
}

/*
	Link to Repl IT
	https://replit.com/@GieorghiiZaporo/USE?v=1
*/
