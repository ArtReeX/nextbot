package console

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"../core"
	"../core/brain"
)

// LaunchingDialog - функция для выполнения последовательности диалога
func LaunchingDialog(network *brain.NeuralNetwork, events chan<- string) {

	// определение считывателя
	reader := bufio.NewReader(os.Stdin)

	for {

		// запрос ввода строки
		fmt.Print("ВЫ      ->: ")

		// считывание строки
		question, error := reader.ReadString('\n')
		if error != nil {
			log.Fatal(error)
		}

		fmt.Print("nextBOT ->: ")
		fmt.Println(core.Input(question, network, events))

		// задержка
		time.Sleep(time.Second * 2)

	}

}
