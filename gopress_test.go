package main

import "testing"

func TestPush(t *testing.T) {
	p := NewGoPress(1) // Kapasitas 1

	// Test 1: Masuk antrean (harus true)
	if ok := p.Push("A"); !ok {
		t.Error("Harusnya A berhasil masuk")
	}

	// Test 2: Antrean penuh (harus false)
	if ok := p.Push("B"); ok {
		t.Error("Harusnya B ditolak")
	}
}
