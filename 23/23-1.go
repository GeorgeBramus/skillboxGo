package main

import (
	"fmt"
	"time"
)

/*
23.1. Сортировка вставками
Напишите функцию, сортирующую массив длины 10 вставками.
*/

const sizeV4 = 10

func sortingWithChoice(a [sizeV4]int) [sizeV4]int {
	for i := 0; i < sizeV4; i++ {
		minIdx := i
		for j := i; j < sizeV4; j++ {
			if a[j] <= a[minIdx] {
				minIdx = j
			}
		}
		a[i], a[minIdx] = a[minIdx], a[i]
	}
	return a
}

func sortingWithInserts(a [sizeV4]int) [sizeV4]int {
	var min int
	for i := 1; i < sizeV4; i++ {
		for j := 0; j < i; j++ {
			if a[i] <= a[j] {
				min = a[i]
				for k := i; k > j; k-- {
					a[k] = a[k-1]
				}
				a[j] = min
			}
		}
	}
	return a
}

func main() {
	a := [sizeV4]int{4, 8, 3, 1, 2, 9, 5, 7, 6}

	fmt.Println("Исходный массив")
	fmt.Println(a)
	fmt.Println()

	fmt.Println("Сортировка вставками")
	start := time.Now()
	fmt.Println(sortingWithInserts(a))
	fmt.Println("Время выполнения", time.Since(start))
	fmt.Println()

	fmt.Println("Сортировка выбором")
	start2 := time.Now()
	fmt.Println(sortingWithChoice(a))
	fmt.Println("Время выполнения", time.Since(start2))
}
