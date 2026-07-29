package main

import "context"

// Semaphore ограничивает параллелизм и поддерживает отменяемое ожидание.
type Semaphore struct {
	slots chan struct{}
}

// New создаёт семафор с заданным числом разрешений.
func New(limit int) *Semaphore {
	if limit <= 0 {
		panic("лимит семафора должен быть больше нуля")
	}

	return &Semaphore{
		slots: make(chan struct{}, limit),
	}
}

// Acquire захватывает разрешение или возвращает ошибку context.
func (s *Semaphore) Acquire(ctx context.Context) error {
	select {
	case s.slots <- struct{}{}:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Release возвращает ранее захваченное разрешение.
func (s *Semaphore) Release() {
	<-s.slots
}
