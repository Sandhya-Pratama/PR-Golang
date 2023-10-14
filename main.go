package main

import "fmt"

func hitungTip(tagihan float64) float64 {
	var tip float64

	if tagihan >= 50 && tagihan <= 300 {
		// Jika tagihan antara 50 dan 300, tip 15%
		tip = 0.15 * tagihan
	} else {
		// Jika tagihan berbeda, tip 20%
		tip = 0.20 * tagihan
	}

	return tip
}

func main() {
	// Nilai tagihan
	tagihan1 := 275.0
	tagihan2 := 430.0
	tagihan3 := 80.00

	// Menghitung tip
	tip1 := hitungTip(tagihan1)
	tip2 := hitungTip(tagihan2)
	tip3 := hitungTip(tagihan3)

	// Menampilkan informasi
	tampilkanInfo := func(tagihan float64, tip float64) {
		total := tagihan + tip
		fmt.Printf("Tagihannya %.2f, tipnya %.2f, dan total nilainya %.2f\n", tagihan, tip, total)
	}

	tampilkanInfo(tagihan1, tip1)
	tampilkanInfo(tagihan2, tip2)
	tampilkanInfo(tagihan3, tip3)
}
