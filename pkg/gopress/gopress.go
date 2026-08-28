package gopress

type GoPress struct {
	queue chan string
}

func New(cap int) *GoPress {
	return &GoPress{
		queue: make(chan string, cap),
	}
}

func (p *GoPress) Push(item string) bool {
	select {
	case p.queue <- item:
		return true
	default:
		return false
	}
}
