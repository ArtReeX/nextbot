package core

import (
	"math/rand"

	"./brain"
)

// Initialize - функция производящая инициализацию нейронной сети
func Initialize() *brain.NeuralNetwork {

	// создание экземпляра нейронной сети
	network := brain.NeuralNetwork{}

	// первоначальное обучение нейронной сети
	FirstTrain(&network)

	return &network

}

// Сompletion - функция предназначенная для завершения сеанса нейронной сети
func Сompletion(network *brain.NeuralNetwork) {

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
