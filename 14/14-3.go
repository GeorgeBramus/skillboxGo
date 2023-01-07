package main

import "fmt"

func changeMultiple(x int) (a int) {
	a = x * 2
	return
}

func changeSum(x int) (a int) {
	a = x + 10
	return
}

func main() {
	var z int = 5
	fmt.Println("Результат:", changeMultiple(changeSum(z)))

}
