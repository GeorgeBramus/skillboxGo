package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	// ***
	// Вводные

	sentences := []string{"First sentence", "Запуск ракеты", "Яркое солнце"}
	chars := []rune{'s', 'р', 'n', 'н', 'т', 'ц'}

	// ***
	// Определим размер двумерного слайса по массиву с предложениями

	// количество предложений - строки
	sLen := len(sentences) // sLen, т.е. "sentences length"

	// количество букв в последнем слове - столбцы
	// с помощью массива определим максимальное количество букв
	maxNumberOfLetters := 0
	for _, sentence := range sentences {
		words := strings.Split(sentence, " ")
		lenLastWord := utf8.RuneCountInString(words[len(words)-1])
		if lenLastWord > maxNumberOfLetters {
			maxNumberOfLetters = lenLastWord
		}
	}
	cLen := maxNumberOfLetters // cLen, т.е. "chars length"

	// ***
	// Создание двумерного слайса

	positions := make([][]int, sLen) // sentences len
	for i := range positions {
		positions[i] = make([]int, cLen) // chars len
	}

	// ***
	// Работа с рунами и запись в двумерный слайс

	for i, sentence := range sentences {
		words := strings.Split(sentence, " ")
		lastWord := words[len(words)-1]
		fmt.Println(lastWord)
		// rLastWord := []rune(lastWord)

		for j, char := range chars {
			for letterIndex, letterInTheWord := range lastWord {
				if letterInTheWord == char {
					positions[i][j] = letterIndex
					fmt.Printf("СОВПАЛО [%v][%v] символ=%v буква=%v %v\n", i, j, string(char), string(letterInTheWord), letterIndex)
				}
			}
		}
	}

	// ***
	// Проверка

	for i := range positions {
		fmt.Println(positions[i])
	}
}
