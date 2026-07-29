package main

import (
	"testing"
	"time"
)

func TestSemaphoreLimitsTotalWeight(t *testing.T) {
	semaphore := New(3)
	semaphore.Acquire(2)

	acquired := make(chan struct{})
	go func() {
		semaphore.Acquire(2)
		close(acquired)
		semaphore.Release(2)
	}()

	select {
	case <-acquired:
		t.Fatal("операция с весом 2 вошла при доступном весе 1")
	case <-time.After(20 * time.Millisecond):
	}

	semaphore.Release(2)

	select {
	case <-acquired:
	case <-time.After(time.Second):
		t.Fatal("операция не получила освободившийся вес")
	}
}

func TestAcquireRejectsWeightAboveCapacity(t *testing.T) {
	semaphore := New(3)

	defer func() {
		if recover() == nil {
			t.Fatal("Acquire не вызвал panic для веса выше ёмкости")
		}
	}()

	semaphore.Acquire(4)
}
