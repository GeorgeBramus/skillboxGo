package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	for i := 1; i < 7; i++ {
		wg.Add(1)
		go func() {
			result := 1
			for j := 1; j <= i; j++ {
				result *= j
			}
			fmt.Println(i, "-", result)
		}()
		wg.Done()
	}
	wg.Wait()
	elapsedTime := time.Since(start)
	time.Sleep(time.Millisecond * 500)
	fmt.Println("Total Time For Execution: " + elapsedTime.String())

	time.Sleep(time.Millisecond * 500)
	fmt.Println("The End")
}
