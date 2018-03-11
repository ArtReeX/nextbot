package core

import (
	"io/ioutil"
	"log"
	"os"
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

// ReadFromFile - функция для загрузки с файла
func ReadFromFile(path string, autoCreate bool) []byte {

	// проверка наличия файла
	CheckExistenceFile(path, autoCreate)

	// чтение файла
	file, error := ioutil.ReadFile(path)
	if error != nil {
		log.Fatal("Error: can not read file [" + path + "]. " + error.Error())
	}

	return file

}

// WriteToFile - функция для загрузки с файла
func WriteToFile(path string, data []byte, overwrite bool, autoCreate bool) {

	// проверка наличия файла
	if CheckExistenceFile(path, autoCreate) && overwrite {

		if error := os.Remove(path); error != nil {
			log.Fatal("Error: can not delete file [" + path + "]. " + error.Error())
		}

		// создание файла
		file, error := os.Create(path)
		if error != nil {
			log.Fatal("Error: сan not create file [" + path + "]. " + error.Error())
		}

		// закрытие файла
		file.Close()

	}

	// открытие файла для сохранения снимка нейронной сети
	file, error := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0777)
	if error != nil {
		log.Fatal("Error: Can not open file [" + path + "]. " + error.Error())
	} else {
		defer file.Close()
	}

	// запись данных в файл
	_, error = file.Write(data)
	if error != nil {
		log.Fatal("Error: сan not write data to file [" + path + "]. " + error.Error())
	}

}

// CheckExistenceFile - функция для проверки существования файла
func CheckExistenceFile(path string, autoCreate bool) bool {

	// проверка наличия файла
	_, error := os.Stat(path)
	if error != nil {

		if autoCreate {

			// создание файла
			file, error := os.Create(path)
			if error != nil {
				log.Fatal("Error: can not create a nonexistent file [" + path + "]. " + error.Error())
			}

			// закрытие файла
			file.Close()

		}

		return false

	}

	return true

}
