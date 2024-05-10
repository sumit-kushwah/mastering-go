package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Dcr decrements the counter for the given key.
func (c *SafeCounter) Dcr(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]--
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func MainMutex() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 5; i++ {
		go c.Inc("somekey")
		go c.Dcr("somekey")
	}
	// fmt.Println(c.Value("somekey"))
	time.Sleep(1 * time.Second)
	fmt.Println(c.Value("somekey"))
}
