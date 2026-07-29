package main

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestAcquireStopsWhenContextExpires(t *testing.T) {
	semaphore := New(1)
	if err := semaphore.Acquire(context.Background()); err != nil {
		t.Fatalf("первый Acquire вернул ошибку: %v", err)
	}
	defer semaphore.Release()

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()

	err := semaphore.Acquire(ctx)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("ожидалась ошибка DeadlineExceeded, получена %v", err)
	}
}

func TestReleaseMakesPermitAvailableAgain(t *testing.T) {
	semaphore := New(1)
	if err := semaphore.Acquire(context.Background()); err != nil {
		t.Fatalf("первый Acquire вернул ошибку: %v", err)
	}
	semaphore.Release()

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()

	if err := semaphore.Acquire(ctx); err != nil {
		t.Fatalf("повторный Acquire вернул ошибку: %v", err)
	}
	semaphore.Release()
}

func TestNewRejectsNonPositiveLimit(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("New не вызвал panic для нулевого лимита")
		}
	}()

	New(0)
}
