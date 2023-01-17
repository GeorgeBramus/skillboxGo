package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

/*
25.1. Написать программу с флагами
Написать программу для нахождения подстроки в кириллической подстроке. Программа должна запускаться с помощью команды:
 go run main.go --str "строка для поиска" --substr "поиска"
Для реализации такой работы с флагами воспользуйтесь пакетом flags, а для поиска подстроки в строке вам понадобятся руны

что нужно сделать
1. Спроектировать алгоритм поиска подстроки.
2. Определить строку и подстроку, используя флаги.
3. Написать алгоритм реализацию для работы со строками UTF-8 (для этого необходимо воспользоваться рунами).
*/

// findSubstrWithContains поиск подстроки с помощью strings.Contains
func findSubstrWithContains(str, substr *string) (bool, error) {
	if str == nil || substr == nil {
		return false, fmt.Errorf("nil pointer")
	}
	return strings.Contains(*str, *substr), nil
}

// findSubstrWithRune поиск подстроки с помощью перебора рун
func findSubstrWithRune(str, substr *string) (bool, error) {
	if str == nil || substr == nil {
		return false, fmt.Errorf("nil pointer")
	}

	rStr := []rune(*str)
	rSubStr := []rune(*substr)

str:
	for i, charStr := range rStr {
		if charStr == rSubStr[0] {
			for j, charSubStr := range rSubStr[1:] {
				if rStr[i+j+1] != charSubStr {
					continue str
				}
			}
			return true, nil
		}
	}

	return false, nil
}

// displayResult показывает результат работы вункции по поиску подстроки
// isFind=true  - если подстрока входит в строку
// isFind=false - подстрока не найдена
// err!=nil     - если приняты флаги со значением nil
func displayResult(isFind bool, err error) {
	if err != nil {
		log.Fatalln("Ошибка приёма данных\nerr:", err)
	} else {
		if isFind {
			fmt.Println("Есть вхождение подстроки.")
		} else {
			fmt.Println("Вхождение не найдено")
		}
	}
}

func main() {
	var str, substr string
	flag.StringVar(&str, "str", "", "usage str")
	flag.StringVar(&substr, "substr", "", "usage substr")
	flag.Parse()

	fmt.Println("Поиск подстроки с помощью strings.Contains")
	isFind, err := findSubstrWithContains(&str, &substr)
	displayResult(isFind, err)

	fmt.Println("\nПоиск подстроки с помощью самописного алгоритма")
	isFind, err = findSubstrWithRune(&str, &substr)
	displayResult(isFind, err)
}
