package core

import (
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
			log.Fatal("Ошибка: невозможно открыть файл со снимком нейронной сети.")
		} else {
			defer file.Close()
		}

		snapshot, error := ioutil.ReadFile(SnapshotFile)
		if error != nil {
			log.Fatal("Ошибка: невозможно прочитать файл со снимком нейронной сети.")
		}

		// загрузка снимка в нейронную сеть
		network.Load(snapshot)

	}

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
			log.Fatal("Ошибка: невозможно удалить файл перед перезаписью снимка нейронной сети.")
		}

	} else {

		// создание файла для сохранения снимка нейронной сети
		file, error := os.Create(SnapshotFile)
		if error != nil {
			log.Fatal("Ошибка: невозможно создать файл для хранения снимка нейронной сети.")
		}

		// закрытие файла
		file.Close()

	}

	// открытие файла для сохранения снимка нейронной сети
	file, error := os.OpenFile(SnapshotFile, os.O_CREATE|os.O_WRONLY, 0777)
	if error != nil {
		log.Fatal("Ошибка: невозможно открыть файл для сохранения снимка нейронной сети.")
	} else {
		defer file.Close()
	}

	// сохранение снимка нейронной сети
	snapshot, error := network.Save()
	if error != nil {
		log.Fatal("Ошибка: невозможно получить снимок нейронной сети.")
	}

	// запись снимка нейронной сети в файл
	_, error = file.Write(snapshot)
	if error != nil {
		log.Fatal("Ошибка: невозможно записать снимок нейронной сети в файл.")
	}

}
