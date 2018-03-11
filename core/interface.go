package core

import (
	"./brain"
)

// Input - функция, предназначенная для запроса боту
func Input(question string, brainBot *brain.NeuralNetwork) string {

	question = FilterTheQuestion(question)

	return question

}
