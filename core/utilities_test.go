package core

import (
	"math/rand"
	"reflect"
	"strconv"
	"testing"
	"time"
)

// ClearText - функция для тестирования очистки текста от лишних символов
func TestClearText(t *testing.T) {

	if ClearText(" Привет, /меня зовут Джон. ' ") != "Привет, меня зовут Джон." {
		t.Error("Error: wrong text cleansing.")
	}

}

// TestWriteAndReadFile - функция для тестирования записи и чтения с файла
func TestWriteAndReadFile(t *testing.T) {

	// инициализация случайных чисел
	rand.Seed(time.Now().UnixNano())

	// создание случайного имени файла
	path := strconv.Itoa(rand.Int()) + ".test"

	// создание структуры для записи
	forWrite := []float64{1.2, 12, 122, 156.7}

	// запись в файл
	WriteToFile(path, Encode(forWrite), true, true)

	// создание структуры для чтения
	forRead := []float64{}

	// чтение с файла
	Decode(ReadFromFile(path, true), &forRead)

	// удаление файла
	RemoveFile(path)

	if !reflect.DeepEqual(forWrite, forRead) {
		t.Error("Error: the structure of the recorded and read object is different.")
	}

}

// TestEncode - функция для тестирования преобразования объектов в JSON формат
func TestEncode(t *testing.T) {

	if string(Encode([]float64{1.2, 12, 122, 156.7})) != "[1.2,12,122,156.7]" {
		t.Error("Error: invalid object encoding.")
	}

}

// TestDecode - функция для тестирования преобразования объектов из JSON формата
func TestDecode(t *testing.T) {

	object := []float64{}

	Decode([]byte("[1.2,12,122,156.7]"), &object)

	if !reflect.DeepEqual(object, []float64{1.2, 12, 122, 156.7}) {
		t.Error("Error: incorrect decoding of an object.")
	}

}

// TestStringToCodeArray - функция тестирования преобразования строки в массив кодов для нейронной сети
func TestStringToCodeArray(t *testing.T) {

	// содание словаря
	dictionary := map[string]float64{
		"Hi,":   1,
		"my":    2,
		"name":  3,
		"John.": 4,
	}

	codeArray := []float64{1, 2, 3, 4, 0, 0, 0, 0, 0, 0}

	if !reflect.DeepEqual(StringToCodeArray("Hi, my name John.", dictionary, 10), codeArray) {
		t.Error("Error: invalid conversion of a string to an array of codes.")
	}

}

// TestCodeArrayToString - функция для тестирования преобразования массива кодов для нейронной сети в строку
func TestCodeArrayToString(t *testing.T) {

	// содание словаря
	dictionary := map[string]float64{
		"Hi,":   1,
		"my":    2,
		"name":  3,
		"John.": 4,
	}

	// создание массива кодов для преобразования
	input := []float64{1, 2, 3, 4, 5, 6, 7, 8}

	// проверка правильности преобразования
	if CodeArrayToString(input, dictionary) != "Hi, my name John." {
		t.Error("Error: incorrect conversion of the code array to a string.")
	}

}
