package main

import "fmt"

func swap(ar []int, i, j int) {
	tmp := ar[i]
	ar[i] = ar[j]
	ar[j] = tmp
}

func bubbleSort(ar []int) []int {
	for i := 0; i < len(ar); i++ {
		for j := len(ar) - 1; j > i; j-- {
			if ar[j-1] > ar[j] {
				swap(ar, j-1, j)
			}
		}
	}
	return ar
}

func main() {
	ar := []int{4, 2, 7, 9, 1, 5, 8, 6, 3}

	fmt.Println("Исходный массив:")
	fmt.Println(ar)
	fmt.Println("Отсортированный массив:")
	fmt.Println(bubbleSort(ar))
}
