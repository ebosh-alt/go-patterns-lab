package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	semaphore := New(3)
	weights := []int64{2, 1, 2, 1}

	var wg sync.WaitGroup
	for taskIndex, weight := range weights {
		id := taskIndex + 1
		wg.Add(1)

		go func() {
			defer wg.Done()

			semaphore.Acquire(weight)
			defer semaphore.Release(weight)

			fmt.Printf("задача %d получила вес %d\n", id, weight)
			time.Sleep(150 * time.Millisecond)
			fmt.Printf("задача %d освободила вес %d\n", id, weight)
		}()
	}

	wg.Wait()
}
