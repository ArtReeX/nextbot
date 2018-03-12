package main

import (
	"sync"
	"time"

	"./console"
	"./core"
)

const (
	// DictionaryFile - файл со словарём
	DictionaryFile = "dictionary.store"
)

// main - функция, являющейся основной
func main() {

	// показ приветствия
	console.ShowGreeting()

	// показ информационного сообщения
	console.ShowInfo()

	// инициализация канала уведомлений
	events := make(chan string)

	// опредение группы контролируемых потоков
	syncGroup := new(sync.WaitGroup)

	// запуск потока уведомлений от нейронной сети
	syncGroup.Add(1)
	go LaunchingEventsScanner(events, syncGroup)

	// инициализация нейронной сети
	network, dictionary := core.Initialize(events)

	// запуск потока диалога
	go console.LaunchingDialog(network, dictionary, events)

	// перевод в режим ожидания окончания всех потоков
	syncGroup.Wait()

	// завершение сеанса нейронной сети
	core.Сompletion(network)

	// показ прощания
	console.ShowFarewell()

	// остановка перед выходом
	time.Sleep(time.Second)

}

// LaunchingEventsScanner - функция обеспечивает постоянное сканирование канала событий
func LaunchingEventsScanner(events <-chan string, syncGroup *sync.WaitGroup) {

	// отложенное завершение потока
	defer syncGroup.Done()

	for {

		// обработка комманд от нейронной сети
		switch <-events {

		case "exit":
			syncGroup.Done()
		case "begin_train":
			console.ShowFirstTrainStart()
		case "end_train":
			console.ShowFirstTrainEnd()

		}

		// переход в режим ожидания на N секунд
		time.Sleep(time.Second)

	}

}
