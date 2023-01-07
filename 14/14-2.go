package main

import (
	"fmt"
	"math/rand"
)

func generatePoint(max int) (x, y int) {
	x = rand.Intn(max)
	y = rand.Intn(max)
	return
}

func changePoints(x, y *int) {
	*x = 2*(*x) + 10
	*y = -3*(*y) - 5
}

func main() {
	var point1x, point1y, point2x, point2y, point3x, point3y, max int

	fmt.Println("Какой будет потолок для координат? (т.е. координаты этой точки будут не больше указанного числа. Этот максимум не должен равняться нулю)")
	fmt.Print("max: ")
	fmt.Scan(&max)

	point1x, point1y = generatePoint(max)
	point2x, point2y = generatePoint(max)
	point3x, point3y = generatePoint(max)

	fmt.Println("Точки до изменения")
	fmt.Printf("[ x1: %v, y1: %v ]\n[ x2: %v, y2: %v ]\n[ x3: %v, y3: %v ]\n", point1x, point1y, point2x, point2y, point3x, point3y)

	changePoints(&point1x, &point1y)
	changePoints(&point2x, &point2y)
	changePoints(&point3x, &point3y)

	fmt.Println("Точки после изменения")
	fmt.Printf("[ x1: %v, y1: %v ]\n[ x2: %v, y2: %v ]\n[ x3: %v, y3: %v ]\n", point1x, point1y, point2x, point2y, point3x, point3y)

}
