package main

import "fmt"

func main() {
	editor := NewEditor("Черновик")
	history := &History{}

	history.Execute(NewReplaceTextCommand(editor, "Пост опубликован"))
	fmt.Println(editor.Text())

	history.Undo()
	fmt.Println(editor.Text())
}
