package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func nowTime() string {
	current := time.Now()
	return fmt.Sprintf("%s", current.Format("2006-01-02 15:04:05"))
}

func main() {
	var (
		inputMessage  string
		numberMessage int
	)

	file, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	sc := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите сообщение: ")
	for sc.Scan() {
		numberMessage++
		inputMessage = sc.Text()
		if sc.Text() == "exit" {
			break
		}
		writer.WriteString(fmt.Sprintf("№%d | %v | %s\n", numberMessage, nowTime(), inputMessage))
		fmt.Print("Ещё сообщение? Или выйти? (exit): ")

	}
	writer.Flush()
	fmt.Println("Файл вроде записался")

}
