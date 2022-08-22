package main

import "fmt"

func main() {
	var dayWeek, numberOfGuests, checkAmount int

	fmt.Println("Введите день недели:")
	fmt.Scan(&dayWeek)
	fmt.Println("Введите число гостей:")
	fmt.Scan(&numberOfGuests)
	fmt.Println("Введите сумму чека:")
	fmt.Scan(&checkAmount)

	discount := 0
	if dayWeek == 1 {
		discount = checkAmount * 10 / 100
		fmt.Println("Скидка по понедельникам:", discount)
	} else if dayWeek == 5 && checkAmount > 10000 {
		discount = checkAmount * 5 / 100
		fmt.Println("Скидка по пятницам:", discount)
	}

	serviceCharge := 0
	if numberOfGuests > 5 {
		serviceCharge = checkAmount * 10 / 100
		fmt.Println("Надбавка за обслуживание:", serviceCharge)
	}

	finalAmount := checkAmount - discount + serviceCharge
	fmt.Println("Сумма к оплате:", finalAmount)
}

/*
	Link to Repl IT
	https://replit.com/@GieorghiiZaporo/Rest?v=1
*/
