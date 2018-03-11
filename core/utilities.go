package core

import (
	"log"
	"regexp"
)

// FilterText - функция, предназначенная для очистки входящего запроса от лишних символов
func FilterText(str string) string {

	reg, error := regexp.Compile("[^ a-z A-Z а-я А-Я 0-9 , ! ? .]")
	if error != nil {
		log.Fatal(error)
	}

	return reg.ReplaceAllString(str, "")

}

// StringArrayToUnique - функция делает матрицу уникальной
func StringArrayToUnique(input []string) []string {
	u := make([]string, 0, len(input))
	m := make(map[string]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}

	return u
}
