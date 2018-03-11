package console

import (
	"fmt"
	"log"
)

// ShowGreeting - функция вывода привествия для пользователя
func ShowGreeting() {

	_, error := fmt.Println(`
		 __________________________________________
		|________________ nextBOT _________________|
		|______________________ on neural networks |
		|__________________________________________|
		|__________________________ Lazarenko A.A. |
		|__________________________________________|
	`)

	if error != nil {
		log.Fatal(error)
	}

}

// ShowInfo - функция вывода прощания для пользователя
func ShowInfo() {

	_, error := fmt.Println(`
		 __________________________________________
		|_________________ INFO: __________________|
		|__________ Write "!bye" to exit. _________|
		|__________________________________________|
	`)

	if error != nil {
		log.Fatal(error)
	}

}

// ShowFarewell - функция вывода прощания для пользователя
func ShowFarewell() {

	_, error := fmt.Println(`
		 __________________________________________
		|________________ GOODBYE :-) _____________|
		|__________________________________________|
		|___ I'll close the window in 1 seconds. __|
		|__________________________________________|
	`)

	if error != nil {
		log.Fatal(error)
	}

}

// ShowFirstTrainStart - функция вывода сообщения о начала первого обучения
func ShowFirstTrainStart() {

	_, error := fmt.Println(`
		 __________________________________________
		|__ Began the first training of the bot. __|
		|__________________________________________|
	`)

	if error != nil {
		log.Fatal(error)
	}

}

// ShowFirstTrainEnd - функция вывода сообщения об окончании первого обучения
func ShowFirstTrainEnd() {

	_, error := fmt.Println(`
		 __________________________________________
		|_____ Finished the first bot training. ___|
		|__________________________________________|
	`)

	if error != nil {
		log.Fatal(error)
	}

}
