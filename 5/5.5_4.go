package main

import (
	"fmt"
)

func main() {
	var amountPayable, firstCoin, secondCoin, thirdCoin uint

	fmt.Println("Укажите сумму для оплаты товара:")
	fmt.Scan(&amountPayable)
	fmt.Println("Укажите номиналы трёх Ваших монет (через пробел):")
	fmt.Scan(&firstCoin, &secondCoin, &thirdCoin)

	possibleAmounts := []uint{
		firstCoin,
		secondCoin,
		thirdCoin,
		firstCoin + secondCoin,
		firstCoin + thirdCoin,
		secondCoin + thirdCoin,
		firstCoin + secondCoin + thirdCoin,
	}

	commonSum := firstCoin + secondCoin + thirdCoin

	if commonSum >= amountPayable {
		var success bool

		for _, amount := range possibleAmounts {
			if amount == amountPayable {
				success = true
				fmt.Println("У Вас есть необходимая сумма для оплаты товара без сдачи.")
				break
			}
		}

		if !success {
			fmt.Println("Вы можете купить товар, но потребуется сдача.")
		}
	} else {
		fmt.Println("У Вас недостаточно средств для покупки товара. Ещё не хватает:", amountPayable-commonSum)
	}

}
