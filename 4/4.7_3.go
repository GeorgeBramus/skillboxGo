package main

import "fmt"

func main() {
	var withdrawalAmount int

	fmt.Println("Банкомат.")
	fmt.Println("Введите сумму снятия со счёта:")
	fmt.Scan(&withdrawalAmount)

	if withdrawalAmount > 100000 {
		fmt.Println("Сумма превышает установленный лимит в 100000.\nЗапросите меньшую сумму.")
	} else if withdrawalAmount%100 != 0 {
		fmt.Println("В банкомате купюры только по 100 рублей.\nСумма должна быть кратной 100.")
	} else {
		fmt.Println("Операция успешно выполнена.\nВы сняли", withdrawalAmount, "рублей.")
	}
}

/*
	Link to Repl IT
	https://replit.com/@GieorghiiZaporo/ATM?v=1
*/
