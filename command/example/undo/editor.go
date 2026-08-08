package main

// Editor хранит редактируемый текст.
type Editor struct {
	text string
}

// NewEditor создаёт редактор с начальным текстом.
func NewEditor(text string) *Editor {
	return &Editor{text: text}
}

// Text возвращает текущий текст.
func (e *Editor) Text() string {
	return e.text
}

func (e *Editor) setText(text string) {
	e.text = text
}

// ReplaceTextCommand заменяет текст и запоминает предыдущее значение.
type ReplaceTextCommand struct {
	editor   *Editor
	text     string
	previous string
}

// NewReplaceTextCommand создаёт команду замены текста.
func NewReplaceTextCommand(editor *Editor, text string) *ReplaceTextCommand {
	return &ReplaceTextCommand{editor: editor, text: text}
}

// Execute заменяет текст, сохраняя значение для отмены.
func (c *ReplaceTextCommand) Execute() {
	c.previous = c.editor.Text()
	c.editor.setText(c.text)
}

// Undo восстанавливает текст, который был до Execute.
func (c *ReplaceTextCommand) Undo() {
	c.editor.setText(c.previous)
}
