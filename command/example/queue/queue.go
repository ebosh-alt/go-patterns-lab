package main

// Command описывает действие, которое можно отложить.
type Command interface {
	Execute()
}

// CommandFunc адаптирует функцию к интерфейсу Command.
type CommandFunc func()

// Execute вызывает обёрнутую функцию.
func (f CommandFunc) Execute() {
	f()
}

// Queue хранит команды до момента выполнения.
type Queue struct {
	commands []Command
}

// Add добавляет команду в конец очереди.
func (q *Queue) Add(command Command) {
	q.commands = append(q.commands, command)
}

// RunAll выполняет команды в порядке добавления и очищает очередь.
func (q *Queue) RunAll() {
	for _, command := range q.commands {
		command.Execute()
	}

	q.commands = nil
}
