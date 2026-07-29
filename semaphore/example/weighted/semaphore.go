package main

import "sync"

// Semaphore ограничивает суммарный вес одновременно выполняемых операций.
type Semaphore struct {
	mu       sync.Mutex
	cond     *sync.Cond
	capacity int64
	used     int64
}

// New создаёт взвешенный семафор с заданной общей ёмкостью.
func New(capacity int64) *Semaphore {
	if capacity <= 0 {
		panic("ёмкость семафора должна быть больше нуля")
	}

	semaphore := &Semaphore{
		capacity: capacity,
	}
	semaphore.cond = sync.NewCond(&semaphore.mu)

	return semaphore
}

// Acquire атомарно захватывает указанный вес или ждёт свободной ёмкости.
func (s *Semaphore) Acquire(weight int64) {
	if weight <= 0 || weight > s.capacity {
		panic("вес должен быть больше нуля и не превышать ёмкость семафора")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	for s.used+weight > s.capacity {
		s.cond.Wait()
	}

	s.used += weight
}

// Release возвращает ранее захваченный вес и будит ожидающие операции.
func (s *Semaphore) Release(weight int64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if weight <= 0 || weight > s.used {
		panic("освобождаемый вес не соответствует захваченному")
	}

	s.used -= weight
	s.cond.Broadcast()
}
