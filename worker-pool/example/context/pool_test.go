package main

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"
)

func TestPoolDrainsTasksAndReportsErrors(t *testing.T) {
	expectedErr := errors.New("задача завершилась с ошибкой")
	pool := New(context.Background(), 2, 4)

	var completed atomic.Int64

	for taskID := range 4 {
		id := taskID

		err := pool.Submit(func(context.Context) error {
			completed.Add(1)
			if id == 2 {
				return expectedErr
			}
			return nil
		})
		if err != nil {
			t.Fatalf("Submit вернул ошибку: %v", err)
		}
	}

	go pool.Close()

	var resultCount int
	var reportedErr error
	for err := range pool.Results() {
		resultCount++
		if err != nil {
			reportedErr = err
		}
	}

	if got := completed.Load(); got != 4 {
		t.Fatalf("выполнено %d задач из 4", got)
	}
	if resultCount != 4 {
		t.Fatalf("получено %d результатов из 4", resultCount)
	}
	if !errors.Is(reportedErr, expectedErr) {
		t.Fatalf("ожидалась ошибка %v, получена %v", expectedErr, reportedErr)
	}
}

func TestPoolCancelsRunningTask(t *testing.T) {
	pool := New(context.Background(), 1, 1)
	started := make(chan struct{})
	stopped := make(chan struct{})

	err := pool.Submit(func(ctx context.Context) error {
		close(started)
		<-ctx.Done()
		close(stopped)
		return ctx.Err()
	})
	if err != nil {
		t.Fatalf("Submit вернул ошибку: %v", err)
	}

	<-started
	pool.Cancel()

	select {
	case <-stopped:
	default:
		t.Fatal("задача не отреагировала на отмену context")
	}
}
