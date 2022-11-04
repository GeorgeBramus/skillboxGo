package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func nowTime() string {
	current := time.Now()
	return fmt.Sprintf("%s", current.Format("2006-01-02 15:04:05"))
}

func main() {
	var (
		fileName      string = "file.txt"
		inputMessage  string
		numberMessage int
	)

	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var b bytes.Buffer

	fmt.Print("Введите сообщение: ")
	for {
		numberMessage++
		fmt.Scan(&inputMessage)
		if inputMessage == "exit" {
			break
		}
		b.WriteString(fmt.Sprintf("№%d | %v | %s\n", numberMessage, nowTime(), inputMessage))
		fmt.Print("Ещё сообщение? Или выйти? (exit): ")
	}

	if err := ioutil.WriteFile(fileName, b.Bytes(), 0666); err != nil {
		panic(err)
	}

	fmt.Println("Файл вроде записался")
}
