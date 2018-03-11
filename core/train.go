package core

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"

	"./brain"
)

const (
	// InitialDialogsFile - файл с начальными диалогами
	InitialDialogsFile = "./train/initial.dialogs"
	// ConversedDialogsFile - файл с конечными диалогами
	ConversedDialogsFile = "./train/conversed.dialogs"
	// DictionaryFile - файл со словарём
	DictionaryFile = "dictionary.store"
)

// FirstTrain - функция для первоначального обучения нейронной сети
func FirstTrain(network *brain.NeuralNetwork) {

	// создание словаря для нейронной сети
	CreateDictionary()

	// установка случайности в нулевое значение
	rand.Seed(0)

	// создание шаблона обучения сети
	patterns := [][][]float64{
		{{0, 0}, {0}},
		{{0, 1}, {1}},
		{{1, 0}, {1}},
		{{1, 1}, {0}},
	}

	// инициализация нейронной сети, структура сети будет содержать 2 входа, 2 скрытых узла и 1 выход
	network.Initialize(2, 2, 1)

	// обучение сети
	network.Train(patterns, 1000, 0.6, 0.4, false)

}

// CreateDictionary - функция используется для создания словаря из существующих диалогов
func CreateDictionary() {

	// проверка существования файла начальных диалогов
	_, error := os.Stat(InitialDialogsFile)
	if error != nil {

		// создание файла для хранения начальных диалогов
		file, error := os.Create(InitialDialogsFile)
		if error != nil {
			log.Fatal("Error: can not create file to store dialogs.")
		}

		// закрытие файла
		file.Close()

	}

	// проверка существования файла словаря
	_, error = os.Stat(DictionaryFile)
	if error == nil {

		// удаление старого файла для хранения диалогов
		error := os.Remove(DictionaryFile)
		if error != nil {
			log.Fatal("Error: unable to delete the dictionary storage file.")
		}

	}

	// создание файла для хранения диалогов
	file, error := os.Create(DictionaryFile)
	if error != nil {
		log.Fatal("Error: unable to create a file to store the dictionary.")
	}

	dictionaryByte, error := ioutil.ReadFile(InitialDialogsFile)
	if error != nil {
		log.Fatal("Error: it is impossible to read the file with the initial dialogs.")
	}

	// закрытие файла
	file.Close()

	// разделение текста по предложениям
	dictionarySentences := strings.Split(string(dictionaryByte), "\r\n")

	// подготовка матрицы для хранения слов
	dictionaryMatrix := make([][]string, len(dictionarySentences))

	// суммарный размер матрицы
	summaryMatrixSize := 0

	// разделение предложений на слова
	for index, element := range dictionarySentences {

		dictionaryMatrix[index] = strings.Split(FilterText(element), " ")
		summaryMatrixSize += len(dictionaryMatrix[index])

	}

	// создание массива со словами
	dictionaryArray := make([]string, summaryMatrixSize)
	dictionaryArrayCount := 0

	// заполнение массива словами
	for indexFirstLayer := range dictionaryMatrix {

		for indexSecondLayer := range dictionaryMatrix[indexFirstLayer] {

			dictionaryArray[dictionaryArrayCount] = dictionaryMatrix[indexFirstLayer][indexSecondLayer]
			dictionaryArrayCount++

		}

	}

	// создание уникального массива
	dictionaryArray = StringArrayToUnique(dictionaryArray)

	// сортировка словаря
	sort.Strings(dictionaryArray)

	// открытие файла для сохранения словаря
	file, error = os.OpenFile(DictionaryFile, os.O_CREATE|os.O_WRONLY, 0777)
	if error != nil {
		log.Fatal("Error: can not open file with dictionary.")
	} else {
		defer file.Close()
	}

	// преобразование словаря в строку для записи
	dictionaryText := ""

	for index, element := range dictionaryArray {

		dictionaryText += "[" + strconv.Itoa(index) + "]" + element + "\r\n"

	}

	// запись словаря в файл
	_, error = file.WriteString(dictionaryText)
	if error != nil {
		log.Fatal("Error: it is impossible to write a file with a dictionary.")
	}

}
