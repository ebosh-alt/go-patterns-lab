package main

import (
	"sync"
	"testing"
	"time"
)

func TestSemaphoreLimitsConcurrentAccess(t *testing.T) {
	semaphore := New(2)
	entered := make(chan struct{}, 3)
	release := make(chan struct{})

	var wg sync.WaitGroup
	for range 3 {
		wg.Add(1)

		go func() {
			defer wg.Done()

			semaphore.Acquire()
			entered <- struct{}{}
			<-release
			semaphore.Release()
		}()
	}

	<-entered
	<-entered

	select {
	case <-entered:
		t.Fatal("третья операция вошла при лимите 2")
	case <-time.After(20 * time.Millisecond):
	}

	close(release)
	wg.Wait()
}

func TestNewRejectsNonPositiveLimit(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("New не вызвал panic для нулевого лимита")
		}
	}()

	New(0)
}
