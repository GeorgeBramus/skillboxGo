package main

import "fmt"

func main() {
	var (
		lastHappyTicket int = 100000
		minCountTickets int
	)
	for i := 100000; i < 1000000; i++ {
		divider := 10
		remainder := []int{}
		numberTicket := i
		for j := 0; j < 6; j++ {
			remainder = append(remainder, numberTicket%divider)
			numberTicket /= divider
		}

		if remainder[0]+remainder[1]+remainder[2] == remainder[3]+remainder[4]+remainder[5] {
			if i-lastHappyTicket > minCountTickets {
				minCountTickets = i - lastHappyTicket
			}
			lastHappyTicket = i
		}
	}
	fmt.Print("Минимальное количество билетов, которые нужно купить, чтобы среди них оказался счастливый: ")
	fmt.Println(minCountTickets)
}
