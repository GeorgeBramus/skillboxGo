package main

import (
	"fmt"
	"io"
)

func lemonadeChange(bills []int) bool {
	var five, ten int
	for _, bill := range bills {
		switch bill {
		case 5:
			five++
		case 10:
			if five > 0 {
				five--
				continue
			} else {
				return false
			}
		case 20:
			if ten > 0 && five > 0 {
				five--
				ten--
				continue
			} else if five >= 3 {
				five -= 3
				continue
			} else {
				return false
			}
		}
	}
	return true
}

func main() {
	var (
		bills []int
		bill  int
	)
	numberClient := 1
	fmt.Println("(!) При вводе нуля для очередного клиента, цикл завершится.")
	for {
		fmt.Printf("Какая купюра у клиента №%d? : ", numberClient)
		_, err := fmt.Scan(&bill)
		if bill > 0 && err != io.EOF {
			bills = append(bills, bill)
		} else {
			break
		}

		numberClient++
	}

	fmt.Println(lemonadeChange(bills))

}
