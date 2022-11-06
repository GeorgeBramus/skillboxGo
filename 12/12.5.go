package main

import "fmt"

func f(dbb, ind, k int, brackets []string) {

	// кладем откр. скобку, только если хватает места
	if dbb <= k-ind-2 {
		brackets[ind] = string('(')
		f(dbb+1, ind+1, k, brackets)
	}

	// закр. скобку можно положить всегда, если dbb > 0
	if dbb > 0 {
		brackets[ind] = string(')')
		f(dbb-1, ind+1, k, brackets)
	}

	// выходим из цикла и печатаем
	if ind == k {
		if dbb == 0 {
			fmt.Println(brackets)
		}
	}
}

func main() {
	var (
		k   int // количество скобок
		dbb int = 0
		ind int = 0
	)

	fmt.Print("Количество пар скоок: ")
	fmt.Scan(&k)
	k = k * 2
	var brackets []string = make([]string, k)

	f(dbb, ind, k, brackets)

}
