package push

import (
	"context"
	"testing"
	"time"

	"github.com/Mulyawan-ts/gopress/pkg/feature/counter"
)

func TestPushWithContext_Timeout(t *testing.T) {
	q := make(chan string, 1)
	c := counter.New()
	p := NewPusher(q, c)

	// Isi antrean sampai penuh
	p.Push("job-1")

	// Buat context dengan timeout 50ms
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	start := time.Now()
	ok := p.PushWithContext(ctx, "job-2")
	elapsed := time.Since(start)

	if ok {
		t.Error("job-2 harusnya gagal karena timeout")
	}

	if elapsed < 50*time.Millisecond {
		t.Errorf("harus menunggu minimal 50ms, tapi cuma %v", elapsed)
	}

	if got := c.Value(); got != 1 {
		t.Errorf("expected dropped count 1, got %d", got)
	}
}
