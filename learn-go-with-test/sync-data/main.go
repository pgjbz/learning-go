package main

import "sync"

type Counter struct {
	mu sync.Mutex
	value uint64
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() uint64 {
	return c.value
}
