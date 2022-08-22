package main

import "fmt"

func clarificationsNeeded(question string) string {
	var answer string

	fmt.Println(question)
	fmt.Print("Попробуем ещё раз? (y/n): ")
	fmt.Scanf("%s\n", &answer)

	return answer
}

func main() {
	var price, discount int

	for {
		fmt.Print("Укажите стоимость товара (руб): ")
		fmt.Scan(&price)
		fmt.Print("Укажите скидку (%): ")
		fmt.Scan(&discount)

		if discount <= 30 {

			discountApplied := float64(price * discount / 100)

			if discountApplied <= 2000 {
				fmt.Printf("%vр. - скидка в самый раз!\n", fmt.Sprintf("%.2f", discountApplied))
				break
			} else {
				answer := clarificationsNeeded("Скидка не должна превышать 2000 рублей.")

				if answer == "y" || answer == "yes" || answer == "" {
					continue
				} else {
					fmt.Println("До свидания!")
					break
				}
			}

		} else {
			answer := clarificationsNeeded("Слишком большая скидка. Скидка не должна превышать 30%.")

			if answer == "y" || answer == "yes" || answer == "" {
				continue
			} else {
				fmt.Println("До свидания!")
				break
			}
		}

	}

}
