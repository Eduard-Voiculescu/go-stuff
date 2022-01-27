package counter

import "sync"

type Counter struct {
	sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.Lock()
	c.value++
	c.Unlock()
}

func (c *Counter) Value() int {
	return c.value
}

func NewCounter() *Counter {
	return &Counter{}
}

/*
Mutex or chanel ?
- Use channels when passing ownership of data
- Use mutexes for managing state
*/
