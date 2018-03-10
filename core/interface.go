package core

// Input - функция, предназначенная для запроса боту
func Input(question string) string {
	question = FilterTheQuestion(question)

	return question
}
