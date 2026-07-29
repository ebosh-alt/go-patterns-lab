package main

import (
	"sync/atomic"
	"testing"
)

func TestPoolExecutesEverySubmittedTask(t *testing.T) {
	const taskCount = 20

	var completed atomic.Int64
	pool := New(3)

	for range taskCount {
		pool.Submit(func() {
			completed.Add(1)
		})
	}

	pool.Close()

	if got := completed.Load(); got != taskCount {
		t.Fatalf("выполнено %d задач из %d", got, taskCount)
	}
}
