package main

import (
	"fmt"
	"math"
)

func main() {
	var overflowUint8, overflowUint16 int
	for i := 0; i < math.MaxUint32; i++ {
		if i > math.MaxUint8 && i > math.MaxUint16 {
			overflowUint8++
			overflowUint16++
		} else if i > math.MaxUint8 {
			overflowUint8++
		}
	}
	message := "Количество переполнений для"
	fmt.Println(message, "uint8\t", overflowUint8)
	fmt.Println(message, "uint16\t", overflowUint16)
}
