package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
)

// sq Функция для возведения каждого числа в квадрат
func sq(nums []int, wg *sync.WaitGroup) chan int {
	out := make(chan int)
	wg.Add(1)
	go func() {
		for _, n := range nums {
			out <- n * n
		}
		close(out)
	}()
	wg.Done()
	return out
}

// mult Функция для умножения каждого числа на 2
func mult(in chan int, wg *sync.WaitGroup) chan int {
	out := make(chan int)
	wg.Add(1)
	go func() {
		for n := range in {
			out <- n * 2
		}
		close(out)
	}()
	wg.Done()
	return out
}

func main() {

	// ***
	// 1. Получение данных от пользователля в бесконечном цикле

	scanner := bufio.NewScanner(os.Stdin)
	var (
		nums []int
		i, n int
		wg   sync.WaitGroup
	)

	for scanner.Scan() {
		if scanner.Err() == io.EOF || scanner.Text() == "стоп" {
			break
		}

		// Передаю в слайс
		n, _ = strconv.Atoi(scanner.Text())
		nums = append(nums, n)

		i++
	}

	// ***
	// 2. Передача данных в функцию для возведения в квадрат
	//    sq(nums...)
	sqChan := make(chan int)
	sqChan = sq(nums, &wg)

	// ***
	// 3. Передача данных в функцию для умножения числа на 2
	//    mult(<-chan int)
	multChan := make(chan int)
	multChan = mult(sqChan, &wg)

	// Горутины закончли работу
	wg.Wait()

	// ***
	// Вывод данных

	fmt.Print("Полученные числа:  ")
	for _, n := range nums {
		fmt.Printf("%d  ", n)
	}
	fmt.Println()

	fmt.Print("Результат возведения в квадрат:  ")
	for value := range sqChan {
		fmt.Printf("%d  ", value)
	}
	fmt.Println()

	fmt.Print("Результат умножения:  ")
	for value := range multChan {
		fmt.Printf("%d  ", value)
	}
	fmt.Println()
}
