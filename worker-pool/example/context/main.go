package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	pool := New(context.Background(), 3, 5)

	for taskID := 1; taskID <= 8; taskID++ {
		id := taskID

		err := pool.Submit(func(ctx context.Context) error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(100 * time.Millisecond):
			}

			if id%4 == 0 {
				return fmt.Errorf("задача %d завершилась с ошибкой", id)
			}

			fmt.Printf("задача %d выполнена\n", id)
			return nil
		})
		if err != nil {
			fmt.Printf("задача %d не принята: %v\n", id, err)
		}
	}

	// Close может ждать освобождения канала результатов, поэтому запускаем его отдельно.
	go pool.Close()

	for err := range pool.Results() {
		if err != nil {
			fmt.Println("ошибка:", err)
		}
	}
}
