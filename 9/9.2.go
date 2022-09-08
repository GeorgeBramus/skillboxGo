package main

import (
	"fmt"
	"math"
)

//goland:noinspection ALL
func main() {
	var firstNumber, secondNumber int16
	fmt.Println("Введите два числа для умножение (через пробел)")
	_, _ = fmt.Scan(&firstNumber, &secondNumber)
	resultMultiple := int32(firstNumber) * int32(secondNumber)

	var answer string
	if resultMultiple > 0 {
		switch {
		case resultMultiple <= math.MaxUint8:
			answer = "uint8"
		case resultMultiple <= math.MaxUint16:
			answer = "uint16"
		default:
			answer = "uint32"
		}
	} else {
		switch {
		case resultMultiple >= math.MinInt8:
			answer = "int8"
		case resultMultiple >= math.MinInt16:
			answer = "int16"
		default:
			answer = "int32"
		}
	}

	fmt.Println(answer)
}
