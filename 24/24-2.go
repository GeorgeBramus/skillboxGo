package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

/*
24.2. Поиск символов в нескольких строках
Напишите функцию, которая на вход принимает массив предложений (длинных строк) и массив символов типа rune, а возвращает 2D-массив, где на позиции [i][j] стоит индекс вхождения символа j из chars в последнее слово в предложении i (строку надо разбить на слова и взять последнее). То есть сигнатура следующая:

func parseTest(sentences []string, chars []rune)
*/

// parseTest создание двумерного слайса
// Запишу в массив, на какой позиции (какой индекс) у соответствующей буквы
// в последнем слове предложения.
// positions[i][j],
// где i - индекс предложения в массиве sentences
// где j - индекс символа в массиве chars
func parseTest(sentences []string, chars []rune) (positions [][]int) {
	// ***
	// Определение размера массива
	// для предания ему физических размеров

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

	positions = make([][]int, sLen) // sentences len
	for i := range positions {
		positions[i] = make([]int, cLen) // chars len
	}

	// ***
	// Работа с рунами и запись в двумерный слайс

	for i, sentence := range sentences {
		words := strings.Split(sentence, " ")
		lastWord := words[len(words)-1]
		rLastWord := []rune(lastWord)
		// bLastWord := []byte(lastWord)

		for j, char := range chars {
			for letterIndex, letterInTheWord := range rLastWord {
				if letterInTheWord == char {
					positions[i][j] = letterIndex
				}
			}
		}
	}

	return
}

func main() {
	sentences := []string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет мама"}
	chars := []rune{'r', 'b', 'L', 'd', 'П', 'М', 'м'}

	positions := parseTest(sentences, chars)

	for i, sentence := range sentences {
		fmt.Println("Для предложения", sentence, "в последнем слове")
		words := strings.Split(sentence, " ")
		lastWord := words[len(words)-1]
		for j, char := range chars {
			if strings.Contains(lastWord, string(char)) {
				fmt.Printf("'%v' position %v\n", string(char), positions[i][j])
			}
		}
		fmt.Println()
	}
}
