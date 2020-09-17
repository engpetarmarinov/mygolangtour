package concurrency

import (
	"github.com/wildalmighty/mygolangtour/utils"
	"log"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]uint64
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string, wg *sync.WaitGroup) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()
	wg.Done()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) uint64 {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return c.v[key]
}

func CountTo(n uint64) {
	defer utils.TimeTrack(time.Now(), "counter")
	c := SafeCounter{v: make(map[string]uint64)}
	var wg = sync.WaitGroup{}
	for i := uint64(0); i < n; i++ {
		wg.Add(1)
		go c.Inc("counter", &wg)
	}
	wg.Wait()
	log.Println("counter value: ", c.Value("counter"))
}
