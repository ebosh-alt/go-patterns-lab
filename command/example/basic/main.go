package main

func main() {
	tv := &TV{}

	onButton := NewButton(NewOnCommand(tv))
	offButton := NewButton(NewOffCommand(tv))

	onButton.Press()
	offButton.Press()
}
