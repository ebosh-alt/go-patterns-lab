package main

// Device - получатель команд включения и выключения.
type Device interface {
	On()
	Off()
}
