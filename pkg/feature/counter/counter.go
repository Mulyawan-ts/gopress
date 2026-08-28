package counter

import "sync/atomic"

type Counter struct {
	value uint64
}

func New() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	atomic.AddUint64(&c.value, 1)
}

func (c *Counter) Value() uint64 {
	return atomic.LoadUint64(&c.value)
}
