package push

import "testing"

func TestPush(t *testing.T) {
	q := make(chan string, 1)
	p := NewPusher(q)

	p.Push("A")

	if p.Push("B") {
		t.Error("B harusnya ditolak karena channel penuh")
	}
}
