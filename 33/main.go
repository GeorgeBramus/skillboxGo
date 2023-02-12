package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ***
// Интерфейс

type Storage interface {
	Add(int) bool
	Size() int
	Values() []int
}

// ***
// Заглушка

type StubStorage struct{}

func (fs *StubStorage) Add(int) bool {
	return true
}

func (fs *StubStorage) Size() int {
	return 3
}

func (fs *StubStorage) Values() []int {
	return []int{1, 2, 3}
}

// ***
// Бизнес-логика

type App struct {
	repository Storage
}

func (a *App) Run() {
	for {
		a.printInfo()
		a.printNumbers()
		if number, ok := a.inputNextNumber(); ok {
			a.storeNumber(number)
		} else {
			break
		}
	}
}

func (a *App) printInfo() {
	msg := "Уникальных чисел в коллекции: %v\n"
	count := a.repository.Size()
	fmt.Printf(msg, count)
}

func (a *App) printNumbers() {
	fmt.Println("Список введённых значений:")
	fmt.Println(strings.Trim(fmt.Sprint(a.repository.Values()), "[]"))
}

func (a *App) inputNextNumber() (int, bool) {
	for {
		fmt.Print("Введите цифру или `end` для завершения: ")
		var input string
		fmt.Scanln(&input)
		if value, err := strconv.Atoi(input); err == nil {
			return value, true
		}
		if input == "end" {
			return 0, false
		}
		fmt.Println("Некорректный ввод")
	}
}

func (a *App) storeNumber(number int) {
	msg := "Число %d присутствует в коллекции\n"
	if ok := a.repository.Add(number); ok {
		msg = "Число %d успешно добавлено\n"
	}
	fmt.Printf(msg, number)
}

// ***
// Подменим заглушку настоящим функционалом

type MemStorage struct {
	numbers []int
}

func NewMemStore() *MemStorage {
	return &MemStorage{
		numbers: make([]int, 0),
	}
}

func (ms *MemStorage) Add(number int) bool {
	if ms.contains(number) {
		return false
	}
	ms.numbers = append(ms.numbers, number)
	return true
}

func (ms *MemStorage) Size() int {
	return len(ms.numbers)
}

func (ms *MemStorage) Values() []int {
	var result []int
	result = append(result, ms.numbers...)
	return result
}

func (ms *MemStorage) contains(number int) bool {
	for _, value := range ms.numbers {
		if value == number {
			return true
		}
	}
	return false
}

func main() {
	// repository := &StubStorage{}
	repository := NewMemStore()
	app := &App{repository}
	app.Run()
}
