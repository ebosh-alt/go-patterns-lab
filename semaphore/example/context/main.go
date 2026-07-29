package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	semaphore := New(1)

	if err := semaphore.Acquire(context.Background()); err != nil {
		panic(err)
	}
	fmt.Println("первая операция получила разрешение")

	ctx, cancel := context.WithTimeout(context.Background(), 80*time.Millisecond)
	defer cancel()

	if err := semaphore.Acquire(ctx); err != nil {
		fmt.Println("вторая операция прекратила ожидание:", err)
	}

	semaphore.Release()
	fmt.Println("первая операция освободила разрешение")

	if err := semaphore.Acquire(context.Background()); err != nil {
		panic(err)
	}
	fmt.Println("третья операция получила освободившееся разрешение")
	semaphore.Release()
}
