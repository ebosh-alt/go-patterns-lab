package main

import (
	"context"
	"sync"
)

// Task описывает работу, которая умеет реагировать на отмену и возвращать ошибку.
type Task func(context.Context) error

// Pool ограничивает параллелизм, хранит очередь задач и управляет их отменой.
type Pool struct {
	tasks     chan Task
	results   chan error
	ctx       context.Context
	cancel    context.CancelFunc
	wg        sync.WaitGroup
	closeOnce sync.Once
}

// New создаёт пул с ограниченной очередью и запускает воркеров.
func New(parent context.Context, workers, queueSize int) *Pool {
	ctx, cancel := context.WithCancel(parent)
	pool := &Pool{
		tasks:   make(chan Task, queueSize),
		results: make(chan error, queueSize+workers),
		ctx:     ctx,
		cancel:  cancel,
	}

	for range workers {
		pool.wg.Add(1)
		go pool.worker()
	}

	return pool
}

// Submit помещает задачу в очередь или возвращает ошибку отменённого context.
func (p *Pool) Submit(task Task) error {
	select {
	case <-p.ctx.Done():
		return p.ctx.Err()
	case p.tasks <- task:
		return nil
	}
}

// Results возвращает канал с результатом каждой выполненной задачи.
func (p *Pool) Results() <-chan error {
	return p.results
}

// Close прекращает приём задач и ждёт обработки всей очереди.
// Результаты нужно читать параллельно с вызовом Close.
func (p *Pool) Close() {
	p.closeOnce.Do(func() {
		close(p.tasks)
		p.wg.Wait()
		p.cancel()
		close(p.results)
	})
}

// Cancel останавливает воркеры и передаёт сигнал отмены выполняющимся задачам.
func (p *Pool) Cancel() {
	p.closeOnce.Do(func() {
		p.cancel()
		p.wg.Wait()
		close(p.results)
	})
}

// worker выполняет задачи до закрытия очереди или отмены context.
func (p *Pool) worker() {
	defer p.wg.Done()

	for {
		select {
		case <-p.ctx.Done():
			return
		case task, ok := <-p.tasks:
			if !ok {
				return
			}

			err := task(p.ctx)

			select {
			case p.results <- err:
			case <-p.ctx.Done():
				return
			}
		}
	}
}
