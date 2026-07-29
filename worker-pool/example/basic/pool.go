package main

import "sync"

// Pool ограничивает количество одновременно выполняемых задач.
type Pool struct {
	tasks chan func()
	wg    sync.WaitGroup
}

// New создаёт пул и запускает заданное количество воркеров.
func New(workers int) *Pool {
	pool := &Pool{
		tasks: make(chan func()),
	}

	for range workers {
		pool.wg.Add(1)
		go pool.worker()
	}

	return pool
}

// Submit передаёт задачу первому свободному воркеру.
func (p *Pool) Submit(task func()) {
	p.tasks <- task
}

// Close прекращает приём задач и ждёт завершения всех воркеров.
func (p *Pool) Close() {
	close(p.tasks)
	p.wg.Wait()
}

// worker выполняет задачи, пока канал не будет закрыт.
func (p *Pool) worker() {
	defer p.wg.Done()

	for task := range p.tasks {
		task()
	}
}
