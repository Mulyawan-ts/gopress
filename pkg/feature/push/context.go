package push

import "context"

// PushWithContext akan menunggu slot antrean kosong sampai context dibatalkan atau timeout
func (p *Pusher) PushWithContext(ctx context.Context, item string) bool {
	select {
	case p.queue <- item:
		return true
	case <-ctx.Done():
		p.counter.Inc()
		return false
	}
}
