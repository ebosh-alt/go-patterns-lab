package main

import (
	"fmt"
	"time"
)

func main() {
	pool := New[int, string](3, 5, func(input int) (string, error) {
		time.Sleep(100 * time.Millisecond)

		if input == 5 {
			return "", fmt.Errorf("значение %d не удалось обработать", input)
		}

		return fmt.Sprintf("число-%d", input*2), nil
	})

	for input := 1; input <= 8; input++ {
		pool.Submit(input)
	}

	go pool.Close()

	for result := range pool.Results() {
		if result.Err != nil {
			fmt.Printf("вход %d: %v\n", result.Input, result.Err)
			continue
		}

		fmt.Printf("вход %d -> %s\n", result.Input, result.Output)
	}
}
