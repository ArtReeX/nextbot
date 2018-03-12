package core

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

// FilterText - функция, предназначенная для очистки текста от лишних символов
func FilterText(str string) string {

	reg, error := regexp.Compile("[^ a-z A-Z а-я А-Я 0-9 , ! ? .]")
	if error != nil {
		log.Fatal(error)
	}

	return reg.ReplaceAllString(str, "")

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

// Encode - функция преобразования объектов в JSON формат
func Encode(object interface{}) []byte {

	objectInJSON, error := json.Marshal(object)
	if error != nil {
		log.Fatal("Error: it is not possible to convert an object into a JSON format. " + error.Error())
	}

	return objectInJSON
}

// Decode - функция преобразования объектов из JSON формата
func Decode(object []byte, structure interface{}) {

	error := json.Unmarshal(object, structure)
	if error != nil {
		log.Fatal("Error: could not convert object from JSON format. " + error.Error())
	}

}

// StringToCodeArray - функция преобразования строки в массив кодов для нейронной сети
func StringToCodeArray(str string, dictionary map[string]float64, nInputs uint) []float64 {

	// разбор входящей строки по словам
	inputString := strings.Split(str, "")

	// создание массива с входящим предложением
	inputArray := make([]float64, NInputs)

	// кодирование слов
	for index, element := range inputString {
		inputArray[index] = dictionary[FilterText(element)]
	}

	return inputArray

}

// CodeArrayToString - функция преобразования массива кодов для нейронной сети в строку
func CodeArrayToString(answerCode []float64, dictionary map[string]float64) string {

	// строка для ответа
	answerString := ""

	// преобразование закодированного ответа в слова
	for _, element := range answerCode {

		for index := range dictionary {

			if dictionary[index] == element {
				answerString += index + " "
			}

		}

	}

	return answerString

}
