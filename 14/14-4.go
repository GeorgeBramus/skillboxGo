package main

import "fmt"

const (
	pi float32 = 3.1415926535
	fi float32 = 1.6180339887
	e  float32 = 2.718281828
)

func plusPi(x float32) float32 {
	return float32(x) + pi
}
func plusFi(x float32) float32 {
	return float32(x) + fi
}
func plusE(x float32) float32 {
	return float32(x) + e
}

func main() {
	var a, b, c float32
	fmt.Println("Прибавим константы к числам, которые сейчас зададим.")
	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)
	fmt.Print("Второе: ")
	fmt.Scan(&b)
	fmt.Print("Третье: ")
	fmt.Scan(&c)
	fmt.Printf("В итоге у нас получилось: 1) %v; 2) %v; 3) %v.\n\n", plusPi(a), plusFi(b), plusE(c))

	fmt.Println("Теперь передадим результат из функции в функцию, по очереди.")
	fmt.Println("Результат:", plusPi(plusFi(plusE(a))))

}
