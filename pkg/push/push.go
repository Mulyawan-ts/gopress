package push

import "github.com/Mulyawan-ts/gopress/pkg/counter"

type Pusher struct {
	queue   chan string
	counter *counter.Counter
}

func NewPusher(q chan string, c *counter.Counter) *Pusher {
	return &Pusher{queue: q, counter: c}
}

func (p *Pusher) Push(item string) bool {
	select {
	case p.queue <- item:
		return true
	default:
		p.counter.Inc()
		return false
	}
}
