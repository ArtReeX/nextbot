package core

import (
	"./brain"
)

// Input - функция, предназначенная для запроса боту
func Input(question string, brainBot *brain.FeedForward) string {
	question = FilterTheQuestion(question)

	return question
}
