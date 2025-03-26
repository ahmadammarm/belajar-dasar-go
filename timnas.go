package main

import "fmt"

func main() {
	poinIndonesia := 6
	poinBahrain := 6

	hasilPertandingan := "Indonesia Menang"

	if hasilPertandingan == "Indonesia Menang" {
		poinIndonesia += 3
	} else {
		poinBahrain += 3
	}

	fmt.Println("Hasil pertandingan dimenangkan oleh Indonesia")
	fmt.Println("Sehingga poin Indonesia menjadi:", poinIndonesia)
}