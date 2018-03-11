package core

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"./brain"
)

const (
	// SnapshotFile - имя файла снимка нейронной сети
	SnapshotFile = "neuralnetwork.snapshot"
)

// Initialize - функция производящая инициализацию нейронной сети
func Initialize(events chan<- string) *brain.NeuralNetwork {

	// создание экземпляра нейронной сети
	network := brain.NeuralNetwork{}

	// проверка существования файла снимка нейронной сети
	_, error := os.Stat(SnapshotFile)
	if error != nil {

		events <- "begin_train"

		// первоначальное обучение нейронной сети
		FirstTrain(&network)

		// сохранение снимка обученной сети в файл
		Сompletion(&network)

		events <- "end_train"

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

	fmt.Println(network.Update([]float64{0, 0}))

	return &network

}

// Сompletion - функция предназначенная для завершения сеанса нейронной сети
func Сompletion(network *brain.NeuralNetwork) {

	// проверка наличия файла снимка нейронной сети
	_, error := os.Stat(SnapshotFile)
	if error == nil {

		// удаление старого снимка нейронной сети
		error := os.Remove(SnapshotFile)
		if error != nil {
			log.Fatal("Error: you can not delete a file for overwriting neural network snapshots.")
		}

	} else {

		// создание файла для сохранения снимка нейронной сети
		file, error := os.Create(SnapshotFile)
		if error != nil {
			log.Fatal("Error: it is impossible to create a file for overwriting neural network snapshots.")
		}

		// закрытие файла
		file.Close()

	}

	// создание файла для сохранения снимка нейронной сети
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
