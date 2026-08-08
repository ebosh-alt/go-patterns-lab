package main

// Command описывает действие, которое можно выполнить позднее.
type Command interface {
	Execute()
}

// OnCommand включает устройство.
type OnCommand struct {
	device Device
}

// NewOnCommand создаёт команду включения устройства.
func NewOnCommand(device Device) Command {
	return OnCommand{device: device}
}

// Execute выполняет включение.
func (c OnCommand) Execute() {
	c.device.On()
}

// OffCommand выключает устройство.
type OffCommand struct {
	device Device
}

// NewOffCommand создаёт команду выключения устройства.
func NewOffCommand(device Device) Command {
	return OffCommand{device: device}
}

// Execute выполняет выключение.
func (c OffCommand) Execute() {
	c.device.Off()
}
