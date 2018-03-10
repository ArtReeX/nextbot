package core

import (
	"fmt"
	"math/rand"

	"./brain"
)

// Initialize - функция, производящая инициализацию нейронной сети
func Initialize() *brain.FeedForward {

	// установка случайности в нулевое значение
	rand.Seed(0)

	// создание шаблона обучения сети
	patterns := [][][]float64{
		{{0, 0}, {0}},
		{{0, 1}, {1}},
		{{1, 0}, {1}},
		{{1, 1}, {0}},
	}

	// создание экземпляра передачи
	brainBot := brain.FeedForward{}

	// инициализация нейронной сети, структура сети будет содержать 2 входа, 2 скрытых узла и 1 выход
	brainBot.Init(2, 2, 1)

	// обучение сети
	brainBot.Train(patterns, 1000, 0.6, 0.4, false)

	// тестирование обученной сети
	fmt.Println(brainBot.Update([]float64{1, 1}))

	return &brainBot
}
