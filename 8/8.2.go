package main

import "fmt"

func main() {
	var day string
	fmt.Println("Дни недели.\nВведите будний день недели: пн, вт, ср, чт, пт:")
	_, _ = fmt.Scan(&day)
	switch day {
	case "пн":
		fmt.Println("понедельник")
		fallthrough
	case "вт":
		fmt.Println("вторник")
		fallthrough
	case "ср":
		fmt.Println("среда")
		fallthrough
	case "чт":
		fmt.Println("четверг")
		fallthrough
	case "пт":
		fmt.Println("пятница")
	default:
		fmt.Println("Это не день недели. Попробуйте снова.")
	}
}
