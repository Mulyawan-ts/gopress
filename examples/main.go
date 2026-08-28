package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Mulyawan-ts/gopress/pkg/gopress"
)

func main() {
	fmt.Println("=== Demo Library GoPress ===")

	// 1. Inisialisasi gopress dengan kapasitas antrean terbatas (3 slot)
	gp := gopress.New(3)

	// 2. Simulasi Push biasa (Instant Drop jika penuh)
	fmt.Println("\n--- 1. Memasukkan Jobs (Instant Push) ---")
	jobs := []string{"job-1", "job-2", "job-3", "job-4"}

	for _, job := range jobs {
		if ok := gp.Push(job); ok {
			fmt.Printf("[SUCCESS] %s berhasil masuk antrean\n", job)
		} else {
			fmt.Printf("[DROPPED] %s ditolak (antrean penuh!)\n", job)
		}
	}

	// 3. Simulasi Push menggunakan Context Timeout
	fmt.Println("\n--- 2. Memasukkan Job dengan Context Timeout ---")
	fmt.Println("Mencoba Push 'job-5' dengan batas waktu tunggu 100ms...")

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	if ok := gp.PushWithContext(ctx, "job-5"); ok {
		fmt.Println("[SUCCESS] job-5 berhasil masuk antrean")
	} else {
		fmt.Println("[TIMEOUT] job-5 gagal masuk setelah menunggu 100ms")
	}

	// 4. Tampilkan total job yang ditolak/dropped
	fmt.Printf("\nTotal job ditolak (Dropped Count): %d\n", gp.DroppedCount())

	// 5. Simulasi Worker yang melakukan Pop data dari antrean
	fmt.Println("\n--- 3. Memprosed Jobs dari Antrean ---")
	for {
		item, ok := gp.Pop()
		if !ok {
			fmt.Println("Antrean sudah kosong.")
			break
		}
		fmt.Printf("[PROCESSING] Memproses %s...\n", item)
		time.Sleep(50 * time.Millisecond) // Simulasi kerja worker
	}
}
