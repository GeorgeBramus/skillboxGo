package main

import "fmt"

func main() {

	a := 1
	b := &a

	c := 2
	d := &c

	e := *b + *d

	a = 3
	c = 7

	fmt.Println(e)

	// s := storage.New()
	// fmt.Println("ещё пустой:", len(s))

	// s.Put(message.New("hello"))
	// fmt.Println("должен быть 1:", len(s))

	// s.Put(message.New("world"))
	// fmt.Println("должен быть 2:", len(s))

	// fmt.Println()
	// for _, m := range s {
	// 	fmt.Println(m)
	// }
}
