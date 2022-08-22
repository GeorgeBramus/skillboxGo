package main

import (
	"fmt"
	"math"
)

func whereStudent(studentNumber, countOfGroups int) {
	row := math.Ceil(float64(studentNumber / countOfGroups))
	groupAnswer := 0
	for i := countOfGroups*int(row) + 1; i < countOfGroups*int(row)+countOfGroups; i++ {
		groupAnswer++
		if studentNumber == i {
			break
		}
	}

	fmt.Println("Студент в группе", groupAnswer)
}

func main() {
	var studentNumber, countOfStudents, countOfGroups int
	fmt.Println("Количество студентов на курсе:")
	fmt.Scan(&countOfStudents)
	fmt.Println("Количество групп:")
	fmt.Scan(&countOfGroups)
	fmt.Println("Порядковый номер студента:")
	fmt.Scan(&studentNumber)

	if studentNumber > countOfStudents {
		fmt.Println("Порядковый номер студента не должен превышать общее количество студентов.")
	} else {
		whereStudent(studentNumber, countOfGroups)
	}
}

/*
	Link to Tepl IT
	https://replit.com/@GieorghiiZaporo/student-search?v=1
*/
