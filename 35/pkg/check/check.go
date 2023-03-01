package check

import (
	"strings"
)

// InputParsing разбор строки из терминала
// которую принимаю от пользователя
func InputParsing(str string) (string, string) {
	studentCharacteristics := strings.Split(str, " ")
	code := studentCharacteristics[0]
	country := studentCharacteristics[1]

	return code, country
}
