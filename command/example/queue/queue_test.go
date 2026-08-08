package main

import (
	"reflect"
	"testing"
)

// TestQueueDefersCommandsUntilRunAll проверяет отложенное выполнение и порядок команд.
func TestQueueDefersCommandsUntilRunAll(t *testing.T) {
	var calls []string
	queue := &Queue{}

	queue.Add(CommandFunc(func() {
		calls = append(calls, "first")
	}))
	queue.Add(CommandFunc(func() {
		calls = append(calls, "second")
	}))

	if len(calls) != 0 {
		t.Fatal("команда выполнилась до RunAll")
	}

	queue.RunAll()

	want := []string{"first", "second"}
	if !reflect.DeepEqual(calls, want) {
		t.Fatalf("вызовы = %v, хотим %v", calls, want)
	}
}

// TestQueueClearsExecutedCommands проверяет, что команда не запускается повторно.
func TestQueueClearsExecutedCommands(t *testing.T) {
	var calls int
	queue := &Queue{}
	queue.Add(CommandFunc(func() { calls++ }))

	queue.RunAll()
	queue.RunAll()

	if calls != 1 {
		t.Fatalf("команда выполнилась %d раз, хотим 1", calls)
	}
}
