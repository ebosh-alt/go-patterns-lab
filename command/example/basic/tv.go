package main

import "fmt"

// TV - конкретный получатель команд.
type TV struct {
	running bool
}

// On включает телевизор.
func (t *TV) On() {
	t.running = true
	fmt.Println("Телевизор включён")
}

// Off выключает телевизор.
func (t *TV) Off() {
	t.running = false
	fmt.Println("Телевизор выключен")
}

// IsRunning возвращает текущее состояние телевизора.
func (t *TV) IsRunning() bool {
	return t.running
}
