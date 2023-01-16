package main

import (
	"fmt"
	"strings"
)

/*
24.2. Поиск символов в нескольких строках
Напишите функцию, которая на вход принимает массив предложений (длинных строк) и массив символов типа rune, а возвращает 2D-массив, где на позиции [i][j] стоит индекс вхождения символа j из chars в последнее слово в предложении i (строку надо разбить на слова и взять последнее). То есть сигнатура следующая:

func parseTest(sentences []string, chars []rune)
*/

func parseTest(sentences []string, chars []rune) (positions [][]int) {
	// взял отдельное предложение
	for i, sentence := range sentences {
		// разбил на слова
		s := strings.Split(sentence, " ")
		// сравниваю руны в последнем слове
		for index, ch := range s[len(s)-1] {
			for j, char := range chars {
				if ch == char {
					// записываю индекс предложения и индекс руны, [i][j]
					positions[i][j] = index // индекс совпавшей руны в последнем слове, index
				} else {
					positions[i][j] = index
				}
			}
		}
	}
	return
}

func main() {
	sentences := []string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}
	chars := []rune{'r', 'b', 'L', 'П', 'М'}

	positions := parseTest(sentences, chars)

	for i, sentence := range positions {
		fmt.Printf("Для предложения \"%v\"\n", sentences[i])
		for j, index := range sentence {
			fmt.Printf("\t'%v' position %v\n", chars[j], index)
		}
	}

	fmt.Println(positions)

	fmt.Println("Конец")
}
