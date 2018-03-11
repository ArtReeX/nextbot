package main

import (
	"sync"
	"time"

	"./console"
	"./core"
)

// main - функция, являющейся основной
func main() {

	// показ приветствия
	console.ShowGreeting()

	// показ информационного сообщения
	console.ShowInfo()

	// инициализация нейронной сети
	network := core.Initialize()

	// опредение группы контролируемых потоков
	syncGroup := new(sync.WaitGroup)

	// запуск потока диалога
	syncGroup.Add(1)
	go console.LaunchingDialog(network, syncGroup)

	// перевод в режим ожидания окончания всех потоков
	syncGroup.Wait()

	// показ прощания
	console.ShowFarewell()

	// завершение сеанса нейронной сети
	core.Сompletion(network)

	// остановка перед выходом
	time.Sleep(time.Second * 2)

}
