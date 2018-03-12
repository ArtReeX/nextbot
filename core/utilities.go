package core

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

// ClearText - функция для очистки текста от лишних символов
func ClearText(str string) string {

	reg, error := regexp.Compile("[^ a-z A-Z а-я А-Я 0-9 , ! ? .]")
	if error != nil {
		log.Fatal(error)
	}

	str = reg.ReplaceAllString(str, "")

	str = strings.TrimSpace(str)

	return str

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

// WriteToFile - функция для чтения с файла
func WriteToFile(path string, data []byte, overwrite bool, autoCreate bool) {

	// проверка наличия файла и создание его, в случае такого параметра
	if CheckExistenceFile(path, autoCreate) && overwrite {

		// удаление файла
		RemoveFile(path)

		// создание файла
		CreareFile(path)

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

// CreareFile - функция для создания файла
func CreareFile(path string) {

	file, error := os.Create(path)
	if error != nil {
		log.Fatal("Error: сan not create file [" + path + "]. " + error.Error())
	}

	// закрытие файла
	file.Close()

}

// RemoveFile - функция для удаления файла
func RemoveFile(path string) {

	if error := os.Remove(path); error != nil {
		log.Fatal("Error: can not delete file [" + path + "]. " + error.Error())
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

// Encode - функция для преобразования объектов в JSON формат
func Encode(object interface{}) []byte {

	objectInJSON, error := json.Marshal(object)
	if error != nil {
		log.Fatal("Error: it is not possible to convert an object into a JSON format. " + error.Error())
	}

	return objectInJSON
}

// Decode - функция для преобразования объектов из JSON формата
func Decode(object []byte, structure interface{}) {

	error := json.Unmarshal(object, structure)
	if error != nil {
		log.Fatal("Error: could not convert object from JSON format. " + error.Error())
	}

}

// StringToCodeArray - функция преобразования строки в массив кодов для нейронной сети
func StringToCodeArray(inputString string, dictionary map[string]float64, nInputs uint) []float64 {

	// разбор входящей строки по словам
	inputArray := strings.Split(ClearText(inputString), " ")

	// создание массива с входящим предложением
	arrayCodes := make([]float64, nInputs)

	// кодирование слов
	for index, element := range inputArray {
		if uint(index) < nInputs {
			arrayCodes[index] = dictionary[element]
		} else {
			break
		}
	}

	return arrayCodes

}

// CodeArrayToString - функция преобразования массива кодов для нейронной сети в строку
func CodeArrayToString(inputArrayWithCodes []float64, dictionary map[string]float64) string {

	// строка для ответа
	answerString := ""

	// преобразование закодированного ответа в слова
	for _, element := range inputArrayWithCodes {

		for index := range dictionary {

			if dictionary[index] == element {
				answerString += index + " "
			}

		}

	}

	return ClearText(answerString)

}
