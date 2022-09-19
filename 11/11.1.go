package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var message string = "Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software."
	words := strings.Split(message, " ")
	var count int8

	for _, word := range words {
		wordRune := []rune(word)
		if unicode.IsUpper(wordRune[0]) {
			count++
		}
	}

	fmt.Println("Количество слов с большой буквы:", count)
}
