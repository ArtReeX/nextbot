package core

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"

	"./brain"
)

const (
	// SnapshotFile - имя файла снимка нейронной сети
	SnapshotFile = "neuralnetwork.snapshot"
)

// Initialize - функция производящая инициализацию нейронной сети
func Initialize() *brain.NeuralNetwork {

	// создание экземпляра нейронной сети
	network := brain.NeuralNetwork{}

	// проверка существования файла снимка нейронной сети
	_, error := os.Stat(SnapshotFile)
	if error != nil {

		// первоначальное обучение нейронной сети
		FirstTrain(&network)

		// сохранение снимка обученной сети в файл
		Сompletion(&network)

	} else {

		// открытие файла для считывания снимка нейронной сети
		file, error := os.OpenFile(SnapshotFile, os.O_CREATE|os.O_RDONLY, 0777)
		if error != nil {
			log.Fatal("Error: it is not possible to open a file to read a snapshot of the neural network.")
		} else {
			defer file.Close()
		}

		snapshot, error := ioutil.ReadFile(SnapshotFile)
		if error != nil {
			log.Fatal("Error: it is impossible to read the neural network snapshot file.")
		}

		// загрузка снимка в нейронную сеть
		network.Load(snapshot)

	}

	// тест
	fmt.Println(network.Update([]float64{0, 0}))

	return &network

}

// Сompletion - функция предназначенная для завершения сеанса нейронной сети
func Сompletion(network *brain.NeuralNetwork) {

	// открытие файла для сохранения снимка нейронной сети
	file, error := os.OpenFile(SnapshotFile, os.O_CREATE|os.O_WRONLY, 0777)
	if error != nil {
		log.Fatal("Error: it is not possible to open a file to save a snapshot of the neural network.")
	} else {
		defer file.Close()
	}

	// сохранение снимка нейронной сети
	snapshot, error := network.Save()
	if error != nil {
		log.Fatal("Error: it is impossible to take a snapshot of the neural network for saving.")
	}

	// запись снимка нейронной сети в файл
	_, error = file.Write(snapshot)
	if error != nil {
		log.Fatal("Error: it is not possible to write to a file to save a snapshot of the neural network.")
	}

}

// FirstTrain - функция для первоначального обучения нейронной сети
func FirstTrain(network *brain.NeuralNetwork) {

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
