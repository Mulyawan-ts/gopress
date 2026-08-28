package gopress

import "testing"

func TestPop(t *testing.T) {
	p := New(1)
	p.Push("A")

	item, ok := p.Pop()

	if !ok || item != "A" {
		t.Errorf("Harusnya mendapatkan 'A', tapi dapat item=%s, ok=%t", item, ok)
	}
}
