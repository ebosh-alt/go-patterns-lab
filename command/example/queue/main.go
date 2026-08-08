package main

import "fmt"

func main() {
	queue := &Queue{}

	queue.Add(CommandFunc(func() {
		fmt.Println("Собрать отчёт")
	}))
	queue.Add(CommandFunc(func() {
		fmt.Println("Отправить отчёт")
	}))

	fmt.Println("Команды добавлены в очередь")
	queue.RunAll()
}
