package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		fileName     string = "hello.txt"
		inputMessage string
	)

	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	err = file.Chmod(0444)
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(file)

	sc := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите сообщение: ")
	for sc.Scan() {
		inputMessage = sc.Text()
		if sc.Text() == "exit" {
			break
		}
		writer.WriteString(fmt.Sprintf("%s\n", inputMessage))
		fmt.Print("Ещё сообщение? Или выйти? (exit): ")

	}
	writer.Flush()
	fmt.Println("Все сообщения записаны")
}
