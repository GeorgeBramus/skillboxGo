package main

import "fmt"

func mergerOfArrays(first [4]int, second [5]int) [9]int {
	var resultArray [len(first) + len(second)]int
	copy(resultArray[:], first[:])
	copy(resultArray[len(first):], second[:])
	return resultArray
}

func main() {
	first := [4]int{1, 2, 3, 4}
	second := [5]int{5, 6, 7, 8, 9}

	n := mergerOfArrays(first, second)

	fmt.Println(n)
}
