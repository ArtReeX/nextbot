package core

// Commands - функция, предназначенная для анализа строки на наличие команд
func Commands(question string, events chan<- string) {

	switch question {

	case "!bye":
		events <- "exit"

	}

}
