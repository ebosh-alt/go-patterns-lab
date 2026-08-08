package main

// UndoableCommand описывает команду, которую можно отменить.
type UndoableCommand interface {
	Execute()
	Undo()
}

// History хранит выполненные команды в порядке их выполнения.
type History struct {
	commands []UndoableCommand
}

// Execute выполняет команду и добавляет её в историю.
func (h *History) Execute(command UndoableCommand) {
	command.Execute()
	h.commands = append(h.commands, command)
}

// Undo отменяет последнюю команду. Пустую историю оставляет без изменений.
func (h *History) Undo() {
	last := len(h.commands) - 1
	if last < 0 {
		return
	}

	h.commands[last].Undo()
	h.commands = h.commands[:last]
}
