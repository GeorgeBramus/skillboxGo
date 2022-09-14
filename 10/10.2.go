package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		depositTimeYears                       int
		sum, monthliInterest, discardedPennies float64
	)
	fmt.Print("Введите сумму: ")
	_, _ = fmt.Scan(&sum)
	fmt.Print("Введите ежемесячный процент: ")
	_, _ = fmt.Scan(&monthliInterest)
	fmt.Print("Количество лет вклада: ")
	_, _ = fmt.Scan(&depositTimeYears)

	for i := 0; i < depositTimeYears*12; i++ {
		newSum := sum + sum*monthliInterest/100
		sum = math.Trunc(newSum*100) / 100
		discardedPennies += newSum - sum
	}

	fmt.Printf("Итоговая сумма: %vр.\nОтброшено: %v", sum, discardedPennies)
}
