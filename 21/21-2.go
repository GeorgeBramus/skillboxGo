package main

import "fmt"

/*
Напишите функцию, которая на вход принимает функцию вида A func (int, int) int, а внутри оборачивает и вызывает её при выходе (через defer).

Вызовите эту функцию с тремя разными анонимными функциями A. Тела функций могут быть любыми, но главное, чтобы все три выполняли разное действие.

*/

func calc(t string, a, b int, A func(int, int) int) {
	defer fmt.Println("результат:", A(a, b))
	fmt.Println()
	fmt.Printf("%v (%v, %v)\n", t, a, b)
}

func main() {
	x := 5
	y := 2
	calc("Сложение", x, y, func(x, y int) int { return x + y })
	calc("Вычитание", x, y, func(x, y int) int { return x - y })
	calc("Умножение", x, y, func(x, y int) int { return x * y })
}
