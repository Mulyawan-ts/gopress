package push

type Pusher struct {
	queue chan string
}

func NewPusher(q chan string) *Pusher {
	return &Pusher{queue: q}
}

func (p *Pusher) Push(item string) bool {
	select {
	case p.queue <- item:
		return true
	default:
		return false
	}
}
