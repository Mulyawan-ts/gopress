package pop

import "testing"

func TestPop(t *testing.T) {
	q := make(chan string, 1)
	q <- "A"

	p := NewPopper(q)
	item, ok := p.Pop()

	if !ok || item != "A" {
		t.Errorf("Harusnya dapat 'A', got item=%s, ok=%t", item, ok)
	}
}
