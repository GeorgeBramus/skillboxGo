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
	resultX, prevX := 1.0, 0.0
	for i := 1; i < 50; i++ {
		if math.Abs(resultX-prevX) < epsilon {
			fmt.Println(resultX)
			break
		}
		prevX = resultX
		resultX += math.Pow(x, float64(i)) / factorial(float64(i))

		fmt.Printf("Итерация: i=%v\n\tprevX=%.4f\n\tresultX=%.4f\n\tepsilon=%v\n", i, prevX, resultX, epsilon)
	}
}
