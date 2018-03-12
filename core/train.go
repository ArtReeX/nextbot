package core

import (
	"math/rand"
	"strings"
	"time"

	"./brain"
)

const (
	// InitialDialogsFile - файл с начальными диалогами
	InitialDialogsFile = "./train/initial.dialogs"
	// EncodedDialogsFile - файл с закодированными диалогами
	EncodedDialogsFile = "./train/encoded.dialogs"
	// DictionaryFile - файл со словарём
	DictionaryFile = "dictionary.store"
)

// FirstTrain - функция для первоначального обучения нейронной сети
func FirstTrain(network *brain.NeuralNetwork) {

	// создание словаря для нейронной сети
	CreateDictionary()

	// кодирование диалогов
	EncodeDialogs()

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

	// считывание начальных диалогов
	dialogsByte := ReadFromFile(InitialDialogsFile, true)

	// подготовка местя для хранения словаря
	dictionaryMap := make(map[string]int)

	// инициализация случайных чисел
	rand.Seed(time.Now().UnixNano())

	// разделение текста по предложениям
	dialogsSentences := strings.Split(string(dialogsByte), "\r\n")

	// разделение предложений на слова
	for _, element := range dialogsSentences {

		for _, element := range strings.Split(FilterText(element), " ") {
			dictionaryMap[FilterText(element)] = rand.Int()
		}

	}

	// запись словаря в файл
	WriteToFile(DictionaryFile, Encode(dictionaryMap), true, true)

}

// EncodeDialogs - функция используется для кодирования диалогов
func EncodeDialogs() {

	// считывание словаря из файла
	var initialDialogs map[string]int
	Decode(ReadFromFile(DictionaryFile, true), &initialDialogs)

	// считывание начальных диалогов
	dialogsByte := ReadFromFile(InitialDialogsFile, true)

	// разделение текста по предложениям
	dialogsSentences := strings.Split(string(dialogsByte), "\r\n")

	// создание матрицы для хранения диалогов
	encodedDialogs := make([][]int, len(dialogsSentences))

	// разделение предложений на слова
	for indexFirstLayer, element := range dialogsSentences {

		sentenses := make([]int, len(strings.Split(FilterText(element), " ")))

		for indexSecondLayer, element := range strings.Split(FilterText(element), " ") {
			sentenses[indexSecondLayer] = initialDialogs[FilterText(element)]
		}

		encodedDialogs[indexFirstLayer] = sentenses

	}

	// запись кодированного диалога в файл
	WriteToFile(EncodedDialogsFile, Encode(encodedDialogs), true, true)

}
