package main

import "fmt"

func main() {
	const size = 10
	var (
		numbers          [size]int
		num              int
		countEvenNumbers int = 0
	)
	for i := 0; i < size; i++ {
		fmt.Printf("Введите элемент №%v: ", i)
		fmt.Scan(&num)
		numbers[i] = num
	}

	for _, number := range numbers {
		if number%2 == 0 {
			countEvenNumbers++
		}
	}

	fmt.Printf("\n\nчётных - %v, нечётных - %v\n", countEvenNumbers, size-countEvenNumbers)
}
