package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Students struct {
	name       string
	age, grade int
}

func newStudent(students *Students, name string, age, grade int) {

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	students := make(map[string]Students)
	i := 0
	for scanner.Scan() {
		if scanner.Err() == io.EOF {
			break
		}

		i++
		studentCharacteristics := strings.Split(scanner.Text(), " ")

		name := studentCharacteristics[0]
		age, _ := strconv.Atoi(studentCharacteristics[1])
		grade, _ := strconv.Atoi(studentCharacteristics[2])

		students[name] = Students{
			name:  name,
			age:   age,
			grade: grade,
		}
	}

	fmt.Println()
	for name, characteristics := range students {
		fmt.Printf("%s\t%v\n", name, characteristics)
	}
}
