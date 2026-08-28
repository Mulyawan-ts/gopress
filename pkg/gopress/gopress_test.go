package gopress

import "testing"

func TestGoPress_Integration(t *testing.T) {
	// 1. Inisialisasi gopress dengan kapasitas 1
	gp := New(1)

	// 2. Push data pertama (harus berhasil)
	if ok := gp.Push("job-1"); !ok {
		t.Error("job-1 harusnya berhasil masuk")
	}

	// 3. Push data kedua (harus ditolak karena buffer penuh)
	if ok := gp.Push("job-2"); ok {
		t.Error("job-2 harusnya ditolak karena buffer penuh")
	}

	// 4. Verifikasi counter dropped bertambah 1
	if got := gp.DroppedCount(); got != 1 {
		t.Errorf("expected dropped count 1, got %d", got)
	}

	// 5. Pop data (harus mendapatkan job-1)
	item, ok := gp.Pop()
	if !ok || item != "job-1" {
		t.Errorf("expected job-1, got item=%s, ok=%t", item, ok)
	}
}
