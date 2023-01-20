package check

import (
	"strconv"
	"strings"
)

// InputParsing разбор строки из терминала
// которую принимаю от пользователя
func InputParsing(str string) (string, int, int) {
	studentCharacteristics := strings.Split(str, " ")
	name := studentCharacteristics[0]
	age, _ := strconv.Atoi(studentCharacteristics[1])
	grade, _ := strconv.Atoi(studentCharacteristics[2])
	return name, age, grade
}
