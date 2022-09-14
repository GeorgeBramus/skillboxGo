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

func newStudent(name string, age, grade int) Students {
	student := Students{
		name:  name,
		age:   age,
		grade: grade,
	}
	return student
}

func stringParsing(str string) (string, int, int) {
	studentCharacteristics := strings.Split(str, " ")
	name := studentCharacteristics[0]
	age, _ := strconv.Atoi(studentCharacteristics[1])
	grade, _ := strconv.Atoi(studentCharacteristics[2])
	return name, age, grade
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
		name, age, grade := stringParsing(scanner.Text())
		students[name] = newStudent(name, age, grade)

	}

	fmt.Println()
	for name, characteristics := range students {
		fmt.Printf("%s\t%v\n", name, characteristics)
	}
}
