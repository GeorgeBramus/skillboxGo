package main

import (
	"fmt"
	"math"
)

func factorial(n float64) float64 {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var x, n float64
	fmt.Print("Введите x: ")
	_, _ = fmt.Scan(&x)
	fmt.Print("Введите точность (знаков после запятой): ")
	_, _ = fmt.Scan(&n)

	epsilon := 1 / math.Pow(10, n)
	prevX := 0.0
	for i := 0; i < 7; i++ {
		x = math.Pow(x, float64(i)) / factorial(float64(i))
		if i > 1 && math.Abs(x-prevX) < epsilon {
			fmt.Println(x)
			break
		}
		prevX = x

		fmt.Printf("Итерация: i=%v\n\tprevX=%.4f\n\tx=%.4f\n\tepsilon=%v\n", i, prevX, x, epsilon)
	}

	//fmt.Println(x)
}
