package main

import "fmt"

/*
24.1. Чётные и нечётные
Напишите функцию, которая принимает массив чисел, а возвращает два массива: один из чётных чисел, второй из нечётных.
*/

func splitArray(a []int) ([]int, []int) {
	var evenElements, oddElements []int
	for _, elem := range a {
		if elem%2 == 0 {
			evenElements = append(evenElements, elem)
		} else {
			oddElements = append(oddElements, elem)
		}
	}
	return evenElements, oddElements
}

func main() {
	a := []int{4, 8, 3, 1, 2, 9, 5, 7, 6, 0}

	fmt.Println("Исходный массив")
	fmt.Println(a)
	fmt.Println()

	eElements, oElements := splitArray(a)
	fmt.Println("Чётные\n", eElements)
	fmt.Println("Нечётные\n", oElements)
}
