package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
var ErrWrongInputParams = errors.New("wrong input params")

type Task func() error

type Limiter struct {
	count int
	limit int
	mx    *sync.Mutex
}

func (c *Limiter) Incr() {
	c.mx.Lock()
	defer c.mx.Unlock()

	c.count++
}

func (c *Limiter) IsLimitExceeded() bool {
	c.mx.Lock()
	defer c.mx.Unlock()

	return c.count >= c.limit
}

// Run starts tasks in N goroutines and stops its work when receiving M errors from tasks.
func Run(tasks []Task, n int, m int) error {
	if len(tasks) == 0 || n <= 0 {
		return ErrWrongInputParams
	}

	tasksCh := make(chan Task, len(tasks))
	quitCh := make(chan struct{}, 1)
	defer close(quitCh)

	errLimiter := &Limiter{
		limit: m,
		mx:    &sync.Mutex{},
	}

	wg := &sync.WaitGroup{}
	wg.Add(n)

	for i := 0; i < n; i++ {
		go worker(tasksCh, quitCh, wg, errLimiter)
	}

	for _, task := range tasks {
		tasksCh <- task
	}
	close(tasksCh)

	wg.Wait()

	quitCh <- struct{}{}

	if errLimiter.IsLimitExceeded() {
		return ErrErrorsLimitExceeded
	}

	return nil
}

func worker(tasksCh <-chan Task, quitCh <-chan struct{}, wg *sync.WaitGroup, errLimiter *Limiter) {
	defer wg.Done()

	for {
		select {
		case <-quitCh:
			return

		default:
		}

		select {
		case task, ok := <-tasksCh:

			if !ok {
				return
			}

			if errLimiter.IsLimitExceeded() {
				return
			}

			if err := task(); err != nil {
				errLimiter.Incr()
			}

		case <-quitCh:

			return
		}
	}
}
