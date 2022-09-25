package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func existFile(name string) bool {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return false
	}
	return true
}

func emptyFile(name string) (int64, bool) {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	if fileInfo.Size() == 0 {
		return 0, true
	} else {
		return fileInfo.Size(), false
	}
}

func main() {
	var fileName string = "file.txt"

	if existFile(fileName) {
		if size, empty := emptyFile(fileName); !empty {

			file, err := os.Open(fileName)
			if err != nil {
				panic(err)
			}
			defer file.Close()

			reader := bufio.NewReader(file)
			buf := make([]byte, size)
			fromFile, err := reader.Read(buf)
			if err != nil && err != io.EOF {
				panic(err)
			}
			fmt.Printf("прочитано %d байт: \n%s", fromFile, buf)

		} else {
			fmt.Printf("Файл %v пуст.\n", fileName)
		}
	} else {
		fmt.Printf("Файл с именем %v не найден.\n", fileName)
	}

}
