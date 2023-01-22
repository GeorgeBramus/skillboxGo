package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
29.2. channels< select operator, graceful shutdown
В работе часто возникает потребность правильно останавливать приложения. Например, когда наш сервер обслуживает соединения, а нам хочется, чтобы все текущие соединения были обработаны и лишь потом произошло выключение сервиса. Для этого существует паттерн graceful shutdown.

Напишите приложение, которое выводит квадраты натуральных чисел на экран, а после получения сигнала ^С обрабатывает этот сигнал, пишет «выхожу из программы» и выходит.
*/

// sq Функция для возведения каждого числа в квадрат
func sq(nums ...int) chan int {
	out := make(chan int)
	go func() {
		// time.Sleep(time.Millisecond * 500)
		for _, n := range nums {
			out <- n * n
		}
	}()
	return out
}

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	nums := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	sqChan := sq(nums...)
	for i := 0; i < len(nums)+10; i++ {
		time.Sleep(time.Millisecond * 250)

		select {
		case <-sigChan:
			fmt.Println("выхожу из программы")
			return
		case sqVal := <-sqChan:
			fmt.Println(sqVal)
		default:
			fmt.Println("ждём")
		}
	}
}
