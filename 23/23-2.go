package main

import (
	"fmt"
)

/*
23.2. Анонимные функции
Напишите анонимную функцию, которая на вход получает массив типа integer, сортирует его пузырьком и переворачивает (либо сразу сортирует в обратном порядке, как посчитаете нужным).
*/

const size = 10

func bubbleSort(a [size]int) func(bool) []int {
	for i := size; i > 0; i-- {
		for j := 1; j < i; j++ {
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]
			}
		}
	}

	return func(reverse bool) []int {
		if reverse {
			var rev []int

			for i := size - 1; i >= 0; i-- {
				rev = append(rev, a[i])
			}
			return rev
		}
		return a[:]
	}
}

func main() {
	a := [size]int{4, 8, 3, 1, 2, 9, 5, 7, 6}

	fmt.Println("Исходный массив")
	fmt.Println(a)
	fmt.Println()

	bubble := bubbleSort(a)
	fmt.Println("Сортировка пузырьком")
	fmt.Println(bubble(false))
	fmt.Println("Перевернём")
	fmt.Println(bubble(true))
}
