package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	d, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		text, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}

		if _, err = d.Write([]byte(text)); err != nil {
			log.Fatalln(err)
		}

		// upperText := []byte{}
		text, err = bufio.NewReader(d).ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(string(text))
	}

}
