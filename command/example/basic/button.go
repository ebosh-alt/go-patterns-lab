package main

// Button - отправитель команды. Он не знает, какое действие будет выполнено.
type Button struct {
	command Command
}

// NewButton создаёт кнопку с назначенной командой.
func NewButton(command Command) Button {
	return Button{command: command}
}

// Press передаёт управление команде.
func (b Button) Press() {
	b.command.Execute()
}
