package gopress

import (
	"context"

	"github.com/Mulyawan-ts/gopress/pkg/feature/counter"
	"github.com/Mulyawan-ts/gopress/pkg/feature/pop"
	"github.com/Mulyawan-ts/gopress/pkg/feature/push"
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

func (g *GoPress) PushWithContext(ctx context.Context, item string) bool {
	return g.pusher.PushWithContext(ctx, item)
}

func (g *GoPress) Pop() (string, bool) {
	return g.popper.Pop()
}

func (g *GoPress) DroppedCount() uint64 {
	return g.counter.Value()
}
