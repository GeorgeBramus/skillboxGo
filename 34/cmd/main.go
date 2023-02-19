package main

import (
	"fmt"
	"time"
)

func main() {
	stop := time.After(time.Second * 2)

	for i := 1; ; i++ {
		select {
		case <-stop:
			fmt.Println("Прошло 2 секунд. Финиш.")
			return
		default:
			time.Sleep(time.Millisecond * 100)
			fmt.Println(i, "Work..")
		}
	}
}
