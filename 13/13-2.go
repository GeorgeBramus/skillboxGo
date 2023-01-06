package main

import "fmt"

func changeVal(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	var a, b int = 10, 20
	fmt.Printf("переменные перед изменением\nпеременная a: %v\nпеременная b: %v\n", a, b)
	changeVal(&a, &b)
	fmt.Printf("переменные после функции изменения\nпеременная a: %v\nпеременная b: %v\n", a, b)
}
