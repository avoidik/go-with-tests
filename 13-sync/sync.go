package synccount

import "sync"

type Counter struct {
	tick int
	mu   sync.Mutex
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.tick++
}

func (c *Counter) Value() int {
	return c.tick
}
