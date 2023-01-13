package main

import (
	"fmt"
	"time"
)

/*
23.2. Анонимные функции
Напишите анонимную функцию, которая на вход получает массив типа integer, сортирует его пузырьком и переворачивает (либо сразу сортирует в обратном порядке, как посчитаете нужным).
*/

const size = 10

func bubbleSort(a [size]int) [size]int {
	for i := size; i > 0; i-- {
		for j := 1; j < i; j++ {
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]
			}
		}
	}
	return a
}

func main() {
	a := [size]int{4, 8, 3, 1, 2, 9, 5, 7, 6}

	fmt.Println("Исходный массив")
	fmt.Println(a)
	fmt.Println()

	fmt.Println("Сортировка пузырьком")
	start := time.Now()
	fmt.Println(bubbleSort(a))
	fmt.Println("Время выполнения", time.Since(start))
}
