package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	semaphore := New(2)

	var wg sync.WaitGroup
	for taskID := 1; taskID <= 8; taskID++ {
		id := taskID
		wg.Add(1)

		go func() {
			defer wg.Done()

			semaphore.Acquire()
			defer semaphore.Release()

			fmt.Printf("задача %d получила разрешение\n", id)
			time.Sleep(150 * time.Millisecond)
			fmt.Printf("задача %d освободила разрешение\n", id)
		}()
	}

	wg.Wait()
}
