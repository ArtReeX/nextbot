package main

import (
	"sync"
	"time"

	"./console"
)

// main - функция, являющейся основной
func main() {
	// показ приветствия
	console.ShowGreeting()

	// показ информационного сообщения
	console.ShowInfo()

	// опредение группы контролируемых потоков
	syncGroup := new(sync.WaitGroup)

	// запуск потока диалога
	syncGroup.Add(1)
	go console.LaunchingDialog(syncGroup)

	// перевод в режим ожидания окончания всех потоков
	syncGroup.Wait()

	// показ прощания
	console.ShowFarewell()

	// остановка перед выходом
	time.Sleep(time.Second * 2)
}
