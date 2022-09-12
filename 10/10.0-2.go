package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Print("Введите через пробел коэффициенты\nи свободный член для квадратного уравнения: ")
	_, _ = fmt.Scan(&a, &b, &c)
	d := math.Pow(b, 2) - 4*a*c
	if d > 0 {
		x1 := (-b + math.Sqrt(d)) / (2 * a)
		x2 := (-b - math.Sqrt(d)) / (2 * a)
		fmt.Printf("x1 = %v, x2 = %v\n", x1, x2)
	} else if d == 0 {
		x := (-b) / (2 * a)
		fmt.Printf("x = %v\n", x)
	} else {
		// Действительная часть
		realPart := (-b) / (2 * a)
		// Мнимая часть
		imaginaryPart := math.Sqrt(math.Abs(d)) / (2 * a)
		x1 := complex(realPart, imaginaryPart)
		x2 := complex(realPart, -imaginaryPart)
		fmt.Printf("x1 = %v, x2 = %v\n", x1, x2)
	}
}
