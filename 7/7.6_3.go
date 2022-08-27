package main

import "fmt"

func main() {
	var treeHeight int
	fmt.Println("Вывод ёлочки.")
	fmt.Print("Введите высоту ёлочки: ")
	fmt.Scan(&treeHeight)

	for stars, spaces, rows := 1, (treeHeight*2-1)/2, 0; rows < treeHeight; stars, spaces, rows = stars+2, spaces-1, rows+1 {
		for i := 0; i < spaces; i++ {
			fmt.Print(" ")
		}
		for i := 0; i < stars; i++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
