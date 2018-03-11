package console

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"../core"
	"../core/brain"
)

// LaunchingDialog - функция для выполнения последовательности диалога
func LaunchingDialog(network *brain.NeuralNetwork, syncGroup *sync.WaitGroup) {

	// отложенное завершение потока
	defer syncGroup.Done()

	// определение считывателя
	reader := bufio.NewReader(os.Stdin)

	for {

		// запрос ввода строки
		fmt.Print("YOU: ")

		// считывание строки
		question, error := reader.ReadString('\n')
		if error != nil {
			log.Fatal(error)
		}

		// определение диалога
		switch core.Commands(question) {
		default:
			fmt.Print("BOT: ")
			fmt.Println(core.Input(question, network))
		case 0:
			return
		}

		// задержка
		time.Sleep(time.Second)

	}

}
