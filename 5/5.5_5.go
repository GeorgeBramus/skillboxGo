package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64

	fmt.Println("Введите коэффициенты a, b и c для вычисления квадратного уравнения.")

	fmt.Print("a: ")
	fmt.Scan(&a)
	fmt.Print("b: ")
	fmt.Scan(&b)
	fmt.Print("c: ")
	fmt.Scan(&c)

	if a != 0 && b == 0 && c == 0 {
		fmt.Println("Неполное квадратное уравнение у которого единственный корень:\nx = 0")
	} else if a != 0 && b == 0 && c != 0 {
		if -c/a > 0 {
			x1 := math.Sqrt(-c / a)
			x2 := -(math.Sqrt(-c / a))
			fmt.Println("У этого уравнения есть два корня:\nx1 =", fmt.Sprintf("%.2f", x1), "\nx2 =", fmt.Sprintf("%.2f", x2))
		} else {
			fmt.Println("Для этого неполного квадратного уравнения нет корней")
		}
	} else if a != 0 && b != 0 && c == 0 {
		x2 := -b / a
		fmt.Println("Это неполное квадратное уравннение имеет два корня:\nx1 = 0\nx2 =", fmt.Sprintf("%.2f", x2))

	} else if a != 0 && b != 0 && c != 0 {
		d := math.Pow(b, 2) - 4*a*c
		if d < 0 {
			fmt.Println("У этого квадратного уравнения нет корней.")
		} else if d == 0 {
			x1 := -b / (2 * a)
			fmt.Println("Квадратное уравнение имеет один действительный корень:\nx =", fmt.Sprintf("%.2f", x1))
		} else {

			x1 := (-b + math.Sqrt(d)) / 2 * a
			x2 := (-b - math.Sqrt(d)) / 2 * a
			fmt.Println("Квадратное уравнение имеет два действительнах корня:\nx1 =", fmt.Sprintf("%.2f", x1), "\nx2 =", fmt.Sprintf("%.2f", x2))
		}
	}

}
