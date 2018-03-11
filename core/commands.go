package core

// Commands - функция, предназначенная для анализа строки на наличие команд
func Commands(question string) int {

	switch FilterTheQuestion(question) {

	case "!BYE":
		return 0
	default:
		return -1

	}

}
