package main

import "testing"

// TestHistoryRestoresPreviousText проверяет отмену последнего изменения.
func TestHistoryRestoresPreviousText(t *testing.T) {
	editor := NewEditor("Черновик")
	history := &History{}

	history.Execute(NewReplaceTextCommand(editor, "Пост опубликован"))
	if got := editor.Text(); got != "Пост опубликован" {
		t.Fatalf("текст = %q, хотим %q", got, "Пост опубликован")
	}

	history.Undo()
	if got := editor.Text(); got != "Черновик" {
		t.Fatalf("текст после Undo = %q, хотим %q", got, "Черновик")
	}
}

// TestHistoryUndoesCommandsInReverseOrder проверяет стековую природу истории.
func TestHistoryUndoesCommandsInReverseOrder(t *testing.T) {
	editor := NewEditor("A")
	history := &History{}

	history.Execute(NewReplaceTextCommand(editor, "B"))
	history.Execute(NewReplaceTextCommand(editor, "C"))
	history.Undo()

	if got := editor.Text(); got != "B" {
		t.Fatalf("текст после первой отмены = %q, хотим %q", got, "B")
	}

	history.Undo()
	if got := editor.Text(); got != "A" {
		t.Fatalf("текст после второй отмены = %q, хотим %q", got, "A")
	}
}
