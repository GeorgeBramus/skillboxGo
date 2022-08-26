package main

import "fmt"

func main() {
	var countHappyTickets int
	for i := 100000; i < 1000000; i++ {
		divider := 10
		remainder := []int{}
		numberTicket := i
		for j := 0; j < 6; j++ {
			remainder = append(remainder, numberTicket%divider)
			numberTicket /= divider
		}
		if remainder[0] == remainder[5] && remainder[1] == remainder[4] && remainder[2] == remainder[3] {
			countHappyTickets++
		}
	}
	fmt.Println("Количество зеркальных билетов среди всех шестизначных номеров от 100000 до 999999:")
	fmt.Println(countHappyTickets)
}
