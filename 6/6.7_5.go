package main

import "fmt"

func main() {
	var firstBasketCapacity, secondBasketCapacity, thirdBasketCapacity int

	fmt.Println("Укажите ёмкость каждой из трёх корзин. Сколько яблок помещается в одну.")

	fmt.Print("Ёмкость первой корзины: ")
	fmt.Scan(&firstBasketCapacity)
	fmt.Print("Ёмкость второй корзины: ")
	fmt.Scan(&secondBasketCapacity)
	fmt.Print("Ёмкость третьей корзины: ")
	fmt.Scan(&thirdBasketCapacity)

	firstBasket, secondBasket, thirdBasket := 0, 0, 0
	sequence := 1

	for {

		if firstBasket == firstBasketCapacity && secondBasket == secondBasketCapacity && thirdBasket == thirdBasketCapacity {
			break
		}

		switch sequence {
		case 1:
			sequence = 2
			if firstBasket < firstBasketCapacity {
				firstBasket++
			} else {
				continue
			}
		case 2:
			sequence = 3
			if secondBasket < secondBasketCapacity {
				secondBasket++
			} else {
				continue
			}
		case 3:
			sequence = 1
			if thirdBasket < thirdBasketCapacity {
				thirdBasket++
			} else {
				continue
			}
		}
	}

	fmt.Println("Все корзины заполнены!")
}
