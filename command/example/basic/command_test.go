package main

import "testing"

// TestButtonsControlTVThroughCommands проверяет, что кнопки управляют ТВ через команды.
func TestButtonsControlTVThroughCommands(t *testing.T) {
	tv := &TV{}
	onButton := NewButton(NewOnCommand(tv))
	offButton := NewButton(NewOffCommand(tv))

	onButton.Press()
	if !tv.IsRunning() {
		t.Fatal("телевизор должен быть включён после команды On")
	}

	offButton.Press()
	if tv.IsRunning() {
		t.Fatal("телевизор должен быть выключен после команды Off")
	}
}
