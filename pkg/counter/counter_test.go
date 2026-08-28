package counter

import "testing"

func TestCounter(t *testing.T) {
	c := New()

	c.Inc() // Tambah 1
	c.Inc() // Tambah 1 lagi

	if got := c.Value(); got != 2 {
		t.Errorf("Harusnya mendapat nilai 2, tapi dapat %d", got)
	}
}
