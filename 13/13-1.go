package main

import "fmt"

func sumOfEvenNumbers(a, b int) int {
	var sum int = 0
	if a > b {
		a, b = b, a
	}
	for i := a; i <= b; i++ {
		if i%2 == 0 {
			sum += i
		}
	}

	return sum
}

func main() {
	var a, b int
	fmt.Println("Задайте диапазон, в котором будут суммироваться все чётные числа")
	fmt.Print("Диапазон будет от: ")
	fmt.Scan(&a)
	fmt.Print("и до: ")
	fmt.Scan(&b)
	fmt.Println()
	fmt.Print("Сумма чётных чисел в указанном диапазоне: ")
	fmt.Print(sumOfEvenNumbers(a, b))
	fmt.Println()
}
