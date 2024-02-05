package concurrency

import (
	"fmt"
	"sync"
)

type Semaphore struct {
	tokens chan struct{}
}

func NewSemaphore(capacity int) *Semaphore {
	return &Semaphore{
		tokens: make(chan struct{}, capacity),
	}
}

func (s *Semaphore) Acquire() {
	s.tokens <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.tokens
}

func (s *Semaphore) Close() {
	close(s.tokens)
}

func TestSemaphores() {
	semaphore := NewSemaphore(2)

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d: Trying to acquire\n", id)
			semaphore.Acquire()
			defer semaphore.Release()
			fmt.Printf("Goroutine %d: Acquired\n", id)
		}(i)
	}

	wg.Wait()
	semaphore.Close()
	fmt.Println("All goroutines finished")
}
