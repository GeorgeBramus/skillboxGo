package main

import "fmt"

func main() {
	var width, height int
	fmt.Println("Шахматная доска")
	fmt.Print("Введите ширину: ")
	fmt.Scan(&width)
	fmt.Print("Введите высоту: ")
	fmt.Scan(&height)

	var evenLine, evenChar bool
	var char string
	for i := 0; i < height; i++ {
		if i%2 == 0 {
			evenLine = true
		} else {
			evenLine = false
		}

		for j := 0; j < width; j++ {
			if j%2 == 0 {
				evenChar = true
			} else {
				evenChar = false
			}

			if evenLine && evenChar || !evenLine && !evenChar {
				char = " "
			} else {
				char = "*"
			}

			fmt.Print(char)
		}
		fmt.Println()
	}
}
