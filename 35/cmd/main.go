package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"code/pkg/check"
	"code/pkg/standart"
	"code/pkg/storage"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	codes := storage.New()
	i := 0
	for scanner.Scan() {
		if scanner.Err() == io.EOF {
			break
		}

		i++
		code, country := check.InputParsing(scanner.Text())
		codes.Put(standart.New(code, country))
	}

	fmt.Println()
	for _, code := range codes {
		fmt.Println(code)
	}
}
