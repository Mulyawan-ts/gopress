package push

import (
	"testing"

	"github.com/Mulyawan-ts/gopress/pkg/counter"
)

func TestPush(t *testing.T) {
	q := make(chan string, 1)
	c := counter.New()
	p := NewPusher(q)

	p.Push("A")

	if p.Push("B") {
		t.Error("B harusnya ditolak karena channel penuh")
	}

	if got := c.Value(); got != 1 {
		t.Errorf("Harusnya dropped count = 1, tapi dapat %d", got)
	}
}
