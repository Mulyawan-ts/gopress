package pop

type Popper struct {
	queue chan string
}

func NewPopper(q chan string) *Popper {
	return &Popper{queue: q}
}

func (p *Popper) Pop() (string, bool) {
	select {
	case item := <-p.queue:
		return item, true
	default:
		return "", false
	}
}
