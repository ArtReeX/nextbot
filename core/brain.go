package core

import (
	"log"
	"strings"

	"./brain"
)

const (
	// SnapshotFile - имя файла снимка нейронной сети
	SnapshotFile = "neuralnetwork.snapshot"
)

// Initialize - функция производящая инициализацию нейронной сети
func Initialize(events chan<- string) (*brain.NeuralNetwork, map[string]float64) {

	// создание экземпляра нейронной сети
	network := brain.NeuralNetwork{}

	// проверка существования файла снимка нейронной сети
	if CheckExistenceFile(SnapshotFile, false) {

		// загрузка снимка в нейронную сеть
		network.Load(ReadFromFile(SnapshotFile, true))

	} else {

		events <- "begin_train"

		// первоначальное обучение нейронной сети
		FirstTrain(&network)

		// сохранение снимка обученной сети в файл
		Сompletion(&network)

		events <- "end_train"

	}

	// загрузка словаря
	var dictionary map[string]float64
	Decode(ReadFromFile(DictionaryFile, false), &dictionary)

	return &network, dictionary

}

// Activate - функция производящая вычисления в нейронной сети
func Activate(str string, network *brain.NeuralNetwork, dictionary map[string]float64) string {

	// разбор входящей строки по словам
	inputString := strings.Split(str, "")

	// создание массива с входящим предложением
	inputArray := make([]float64, NInputs)

	// кодирование слов
	for index, element := range inputString {
		inputArray[index] = dictionary[FilterText(element)]
	}

	// получение ответа от нейронной сети
	answerCode := network.Update(inputArray)

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

// Сompletion - функция предназначенная для завершения сеанса нейронной сети
func Сompletion(network *brain.NeuralNetwork) {

	// сохранение снимка нейронной сети
	snapshot, error := network.Save()
	if error != nil {
		log.Fatal("Error: it is impossible to take a snapshot of the neural network for saving.")
	} else {
		WriteToFile(SnapshotFile, snapshot, true, true)
	}

}
