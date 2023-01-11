package main

import (
	"fmt"
	"math"
)

/*
20.1 Расчёт по формуле
Напишите функцию, производящую следующие вычисления.
S = 2 × x + y ^ 2 − 3/z, где x — int16, y — uint8, a z — float32.
Тип S должен быть во float32.
*/

func calculation(x int16, y uint16, z float32) float32 {
	c := float64(y)
	return 2*float32(x) + float32(math.Pow(c, 2)) - 3/z
}

func main() {
	var (
		a int16   = 2
		b uint16  = 8
		c float32 = 12.345
	)

	s := calculation(a, b, c)

	fmt.Println(s)
}
