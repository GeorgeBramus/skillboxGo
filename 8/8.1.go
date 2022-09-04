package main

import "fmt"

func main() {
	var month, season string
	fmt.Println("Времена года.\nВведите месяц:")
	fmt.Scan(&month)
	season = "Время года -"
	switch month {
	case "декабрь", "январь", "февраль":
		fmt.Println(season, "зима")
	case "март", "апрель", "май":
		fmt.Println(season, "весна")
	case "июнь", "июль", "август":
		fmt.Println(season, "лето")
	case "сентябрь", "октябрь", "ноябрь":
		fmt.Println(season, "осень")
	default:
		fmt.Println("Это точно месяц? Попробуйте ещё раз.")
	}
}
