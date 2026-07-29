package main

import (
	"fmt"
	"time"
)

func main() {
	pool := New(3)

	for taskID := 1; taskID <= 8; taskID++ {
		id := taskID

		pool.Submit(func() {
			fmt.Printf("задача %d началась\n", id)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("задача %d завершилась\n", id)
		})
	}

	pool.Close()
}
