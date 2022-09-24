package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var numbers []int
	sourceStr := "a10 10 20b 20 30c30 30 dd"
	words := strings.Split(sourceStr, " ")
	for _, word := range words {
		number, err := strconv.Atoi(word)
		if err == nil {
			numbers = append(numbers, number)
		}
	}
	fmt.Println("Исходная строка:\n", sourceStr)
	fmt.Println("В строке содержатся числа в десятичном формате:\n", numbers)
}
