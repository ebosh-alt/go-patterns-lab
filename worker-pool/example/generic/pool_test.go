package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestPoolReturnsTypedResultsAndErrors(t *testing.T) {
	expectedErr := errors.New("неподдерживаемое значение")

	pool := New[int, string](3, 4, func(input int) (string, error) {
		if input == 3 {
			return "", expectedErr
		}
		return fmt.Sprintf("число-%d", input*2), nil
	})

	for input := 1; input <= 4; input++ {
		pool.Submit(input)
	}

	go pool.Close()

	results := make(map[int]Result[int, string])
	for result := range pool.Results() {
		results[result.Input] = result
	}

	if len(results) != 4 {
		t.Fatalf("получено %d результатов из 4", len(results))
	}
	if got := results[2].Output; got != "число-4" {
		t.Fatalf("для входа 2 ожидалось число-4, получено %q", got)
	}
	if !errors.Is(results[3].Err, expectedErr) {
		t.Fatalf("ожидалась ошибка %v, получена %v", expectedErr, results[3].Err)
	}
}
