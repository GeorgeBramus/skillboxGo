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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	students := make(map[int]struct{})
	i := 0
	for scanner.Scan() {
		if scanner.Err() == io.EOF {
			break
		}

		i++
		studentCharacteristics := strings.Split(scanner.Text(), " ")

		age, _ := strconv.Atoi(studentCharacteristics[1])
		grade, _ := strconv.Atoi(studentCharacteristics[2])

		students[i] = Students{
			name:  studentCharacteristics[0],
			age:   age,
			grade: grade,
		}
	}

	fmt.Println(students)
}
