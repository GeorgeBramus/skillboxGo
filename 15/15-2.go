package main

import "fmt"

const size = 5

func reverseArray(numbers [size]int) [size]int {
	var newNumbers [size]int
	for i, number := range numbers {
		newNumbers[size-i-1] = number
	}
	return newNumbers
}

func main() {
	var (
		numbers, newNumbers [size]int
		num                 int
	)
	fmt.Printf("Введите %v чисел для массива\n", size)
	for i := 0; i < size; i++ {
		fmt.Printf("Введите элемент №%v: ", i)
		fmt.Scan(&num)
		numbers[i] = num
	}

	fmt.Println("\nМассив, который получился")
	for _, number := range numbers {
		fmt.Printf("%v ", number)
	}

	fmt.Println("\n\nИ развернём(реверсируем) массив")
	newNumbers = reverseArray(numbers)
	for _, number := range newNumbers {
		fmt.Printf("%v ", number)
	}
	fmt.Println()
}
