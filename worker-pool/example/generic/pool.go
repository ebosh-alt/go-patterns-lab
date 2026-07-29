package main

import "sync"

// Process преобразует входное значение в результат или ошибку.
type Process[I, O any] func(I) (O, error)

// Result связывает исходные данные с результатом их обработки.
type Result[I, O any] struct {
	Input  I
	Output O
	Err    error
}

// Pool выполняет однотипную функцию над входными значениями параллельно.
type Pool[I, O any] struct {
	tasks   chan I
	results chan Result[I, O]
	process Process[I, O]
	wg      sync.WaitGroup
}

// New создаёт типизированный пул и запускает воркеров.
func New[I, O any](workers, queueSize int, process Process[I, O]) *Pool[I, O] {
	pool := &Pool[I, O]{
		tasks:   make(chan I, queueSize),
		results: make(chan Result[I, O], queueSize+workers),
		process: process,
	}

	for range workers {
		pool.wg.Add(1)
		go pool.worker()
	}

	return pool
}

// Submit добавляет входное значение в очередь.
func (p *Pool[I, O]) Submit(input I) {
	p.tasks <- input
}

// Results возвращает доступный только для чтения канал результатов.
func (p *Pool[I, O]) Results() <-chan Result[I, O] {
	return p.results
}

// Close закрывает очередь и канал результатов после выхода всех воркеров.
// Результаты нужно читать параллельно с вызовом Close.
func (p *Pool[I, O]) Close() {
	close(p.tasks)
	p.wg.Wait()
	close(p.results)
}

// worker применяет общую функцию обработки к каждому входному значению.
func (p *Pool[I, O]) worker() {
	defer p.wg.Done()

	for input := range p.tasks {
		output, err := p.process(input)
		p.results <- Result[I, O]{
			Input:  input,
			Output: output,
			Err:    err,
		}
	}
}
