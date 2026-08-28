package gopress

import (
	"github.com/Mulyawan-ts/gopress/pkg/counter"
	"github.com/Mulyawan-ts/gopress/pkg/pop"
	"github.com/Mulyawan-ts/gopress/pkg/push"
)

type GoPress struct {
	pusher  *push.Pusher
	popper  *pop.Popper
	counter *counter.Counter
}

func New(cap int) *GoPress {
	q := make(chan string, cap)
	c := counter.New()
	return &GoPress{
		pusher:  push.NewPusher(q, c),
		popper:  pop.NewPopper(q),
		counter: c,
	}
}

func (g *GoPress) Push(item string) bool {
	return g.pusher.Push(item)
}

func (g *GoPress) Pop() (string, bool) {
	return g.popper.Pop()
}

func (g *GoPress) DroppedCount() uint64 {
	return g.counter.Value()
}
