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

	// NInputs - количество входящих нейронов
	NInputs = 50
	// NHiddens - количество скрытых нейронов
	NHiddens = 50
	// NOutputs - количество исходящих нейронов
	NOutputs = 50
)

// FirstTrain - функция для первоначального обучения нейронной сети
func FirstTrain(network *brain.NeuralNetwork) {

	// установка случайности в нулевое значение
	rand.Seed(0)

	// инициализация нейронной сети, структура сети будет содержать 2 входа, 2 скрытых узла и 1 выход
	network.Initialize(NInputs, NHiddens, NOutputs)

	// создание словаря для нейронной сети
	CreateDictionary()

	// кодирование диалогов
	EncodeDialogs()

	// создание шеблона нейронной сети
	CreatePatternsForTrain(NInputs, NOutputs)

	// обучение сети
	network.Train(CreatePatternsForTrain(NInputs, NOutputs), 500, 0.6, 0.4, false)

}

// CreatePatternsForTrain - функция используется для создания шаблона обучения
func CreatePatternsForTrain(nInputs, nOutputs uint) [][][]float64 {

	// считывание закодированых диалогов из файла
	var dialogs [][]float64
	Decode(ReadFromFile(EncodedDialogsFile, false), &dialogs)

	patterns := make([][][]float64, len(dialogs)/2)
	indexDialogs := 0

	// создание шаблона с промежуточными нулями
	for indexPatterns := 0; indexPatterns < len(patterns); indexPatterns++ {

		patterns[indexPatterns] = make([][]float64, 2)

		patterns[indexPatterns][0] = make([]float64, nInputs)
		for index, element := range dialogs[indexDialogs] {
			patterns[indexPatterns][0][index] = element
		}

		indexDialogs++

		patterns[indexPatterns][1] = make([]float64, nOutputs)
		for index, element := range dialogs[indexDialogs] {
			patterns[indexPatterns][1][index] = element
		}

		indexDialogs++

	}

	return patterns

}

// CreateDictionary - функция используется для создания словаря из существующих диалогов
func CreateDictionary() {

	// считывание начальных диалогов
	dialogsByte := ReadFromFile(InitialDialogsFile, true)

	// подготовка местя для хранения словаря
	dictionaryMap := make(map[string]float64)

	// инициализация случайных чисел
	rand.Seed(time.Now().UnixNano())

	// разделение текста по предложениям
	dialogsSentences := strings.Split(string(dialogsByte), "\r\n")

	// разделение предложений на слова
	for _, element := range dialogsSentences {

		for _, element := range strings.Split(ClearText(element), " ") {
			dictionaryMap[ClearText(element)] = float64(rand.Int63())
		}

	}

	// запись словаря в файл
	WriteToFile(DictionaryFile, Encode(dictionaryMap), true, true)

}

// EncodeDialogs - функция используется для кодирования диалогов
func EncodeDialogs() {

	// считывание словаря из файла
	var initialDialogs map[string]float64
	Decode(ReadFromFile(DictionaryFile, true), &initialDialogs)

	// считывание начальных диалогов
	dialogsByte := ReadFromFile(InitialDialogsFile, true)

	// разделение текста по предложениям
	dialogsSentences := strings.Split(string(dialogsByte), "\r\n")

	// создание матрицы для хранения диалогов
	encodedDialogs := make([][]float64, len(dialogsSentences))

	// разделение предложений на слова
	for indexFirstLayer, element := range dialogsSentences {

		sentenses := make([]float64, len(strings.Split(ClearText(element), " ")))

		for indexSecondLayer, element := range strings.Split(ClearText(element), " ") {
			sentenses[indexSecondLayer] = initialDialogs[ClearText(element)]
		}

		encodedDialogs[indexFirstLayer] = sentenses

	}

	// запись кодированного диалога в файл
	WriteToFile(EncodedDialogsFile, Encode(encodedDialogs), true, true)

}
