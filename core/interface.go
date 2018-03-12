package core

import (
	"./brain"
)

// Input - функция, предназначенная для запроса боту
func Input(question string, network *brain.NeuralNetwork, dictionary map[string]float64, events chan<- string) string {

	question = ClearText(question)

	// передача строки в функцию проверки комманд
	Commands(question, events)

	return Activate(question, network, dictionary)

}
