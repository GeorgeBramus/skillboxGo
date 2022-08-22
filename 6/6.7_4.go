package main

import "fmt"

func main() {
	first, second, third := 0, 0, 0
	for {
		if first < 10 {
			first++
		}
		if second < 100 {
			second++
		}
		if third < 1000 {
			third++
		} else {
			break
		}
	}
	fmt.Printf("first: %v, second: %v, third: %v\n", first, second, third)
}
