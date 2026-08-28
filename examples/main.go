package main

import (
	"fmt"
	"time"

	"github.com/Mulyawan-ts/gopress/pkg/gopress"
)

func main() {
	fmt.Println("=== Demo Library GoPress ===")

	// 1. Inisialisasi gopress dengan kapasitas antrean terbatas (3 slot)
	gp := gopress.New(3)

	// 2. Simulasi Push beberapa job (melebihi kapasitas)
	jobs := []string{"job-1", "job-2", "job-3", "job-4", "job-5"}

	fmt.Println("\n--- Memasukkan Jobs ---")
	for _, job := range jobs {
		ok := gp.Push(job)
		if ok {
			fmt.Printf("[SUCCESS] %s berhasil masuk ke antrean\n", job)
		} else {
			fmt.Printf("[DROPPED] %s ditolak (antrean penuh!)\n", job)
		}
	}

	// 3. Tampilkan total job yang ditolak
	fmt.Printf("\nTotal job ditolak (Backpressure): %d\n", gp.DroppedCount())

	// 4. Simulasi Worker yang melakukan Pop data dari antrean
	fmt.Println("\n--- Memproses Jobs dari Antrean ---")
	for {
		item, ok := gp.Pop()
		if !ok {
			fmt.Println("Antrean sudah kosong.")
			break
		}
		fmt.Printf("[PROCESSING] Memproses %s...\n", item)
		time.Sleep(100 * time.Millisecond) // Simulasi kerja worker
	}

	// 5. Cek ulang antrean setelah diproses
	_, ok := gp.Pop()
	fmt.Printf("\nStatus Pop akhir (apakah ada data?): %t\n", ok)
}
