package main

// Semaphore ограничивает количество одновременных владельцев разрешения.
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

// Acquire захватывает разрешение или ждёт освобождения места.
func (s *Semaphore) Acquire() {
	s.slots <- struct{}{}
}

// Release возвращает ранее захваченное разрешение.
func (s *Semaphore) Release() {
	<-s.slots
}
