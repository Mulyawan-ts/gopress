package push

import "testing"

func TestPush(t *testing.T) {
	p := New(1)
	p.Push("A")

	if p.Push("B") {
		t.Error("B ditolak karena antrean full")
	}
}
