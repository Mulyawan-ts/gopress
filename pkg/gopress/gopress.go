package gopress

import (
	"github.com/Mulyawan-ts/gopress/pkg/pop"
	"github.com/Mulyawan-ts/gopress/pkg/push"
)

type GoPress struct {
	pusher *push.Pusher
	popper *pop.Popper
}

func New(cap int) *GoPress {
	q := make(chan string, cap)
	return &GoPress{
		pusher: push.NewPusher(q),
		popper: pop.NewPopper(q),
	}
}

func (g *GoPress) Push(item string) bool {
	return g.pusher.Push(item)
}

func (g *GoPress) Pop() (string, bool) {
	return g.popper.Pop()
}
