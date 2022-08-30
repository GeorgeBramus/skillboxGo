package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var (
		count              int
		thirdFile          bool
		messageAboutRecord string
	)

	flag.IntVar(&count, "n", 1, "number of lines to read from the file")
	flag.Parse()

	files := flag.Args()
	str := ""
	if len(files) > 2 {
		thirdFile = true
	}
	for i, file := range files {
		if i < 2 {
			f, err := os.Open(file)
			if err != nil {
				log.Printf("Error opening file (%v). err: %v", file, err)
				os.Exit(1)
			}

			buf := bufio.NewScanner(f)
			for buf.Scan() {
				str += buf.Text() + "\n"
			}
			f.Close()
		} else {
			f, err := os.Create(file)
			if err != nil {
				log.Printf("Error creating %v file. Err: %v\n", f, err)
			}

			w := bufio.NewWriter(f)
			n, err := w.WriteString(str)
			if err != nil {
				log.Printf("Error writing to file %v. Err: %v\n", f, err)
			}
			messageAboutRecord = fmt.Sprintf("Успешно записано %d байт в новый файл %v\n", n, file)

			w.Flush()
			f.Close()
		}
	}

	if !thirdFile {
		fmt.Println(str)
	} else {
		fmt.Println(messageAboutRecord)
	}
}
