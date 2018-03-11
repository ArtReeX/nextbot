package core

import (
	"./brain"
)

// Input - функция, предназначенная для запроса боту
func Input(question string, brainBot *brain.NeuralNetwork, events chan<- string) string {

	question = FilterText(question)

	// передача строки в функцию проверки комманд
	Commands(question, events)

	return question

}
