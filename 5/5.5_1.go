package main

import "fmt"

func main() {
	var (
		x, y                  int32
		coordinatePlaneNumber string
		justTheWordQuarter    string = " четверти"
	)

	fmt.Println("Введите X и Y через пробел:")
	fmt.Scan(&x, &y)

	if x == 0 || y == 0 {
		coordinatePlaneNumber = " начале"
		justTheWordQuarter = ""
	} else if x >= 0 && y >= 0 {
		coordinatePlaneNumber = " I"
	} else if x <= 0 && y >= 0 {
		coordinatePlaneNumber = "о II"
	} else if x <= 0 && y <= 0 {
		coordinatePlaneNumber = " III"
	} else if x >= 0 && y <= 0 {
		coordinatePlaneNumber = " IV"
	}

	fmt.Printf("Точка находится в%s%s координатной плоскости\n", coordinatePlaneNumber, justTheWordQuarter)
}
