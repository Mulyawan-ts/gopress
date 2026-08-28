# GoPress 🚀

[![Go Reference](https://pkg.go.dev/badge/github.com/Mulyawan-ts/gopress.svg)](https://pkg.go.dev/github.com/Mulyawan-ts/gopress)
[![Go Report Card](https://goreportcard.com/badge/github.com/Mulyawan-ts/gopress)](https://goreportcard.com/report/github.com/Mulyawan-ts/gopress)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**GoPress** adalah library antrean (*in-memory queue*) berkinerja tinggi untuk Go. Library ini dirancang dengan mekanisme *backpressure* otomatis (menolak job baru saat antrean penuh) dan pencatatan statistik secara *thread-safe* serta *zero-allocation*.

---

## 🌟 Fitur Utama

* **Modul Terisolasi**: Dibangun dengan arsitektur fitur yang modular dan mudah diuji.
* **Non-Blocking Backpressure**: Otomatis melempar item yang melebih kapasitas tanpa menghentikan eksekusi (*drop on full*).
* **Thread-Safe Metrics**: Menggunakan `sync/atomic` untuk mencatat jumlah *dropped jobs* secara akurat.
* **Zero Allocation**: Dioptimalkan untuk kecepatan tinggi dengan beban alokasi memori $0$ byte per operasi.

---

## 📦 Instalasi

Tambahkan `gopress` ke dalam proyek Go Anda:

```bash
go get [github.com/Mulyawan-ts/gopress](https://github.com/Mulyawan-ts/gopress)
