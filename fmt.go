package main

import "fmt"

// package format dalam go adalah package yang digunakan untuk memformat string, angka, dan tipe data lainnya.
// Package ini menyediakan fungsi-fungsi untuk memformat string dengan cara yang lebih fleksibel dan mudah dibaca.

// contoh function dari format ini antara lain :
// 1. Sprintf: untuk memformat string dengan cara yang lebih fleksibel dan mudah dibaca
// 2. Printf: untuk mencetak string ke layar dengan format yang ditentukan
// 3. Fprintf: untuk mencetak string ke file dengan format yang ditentukan
// 4. Scanf: untuk membaca input dari pengguna dengan format yang ditentukan
// 5. Println: untuk mencetak string ke layar tanpa format yang ditentukan

func deteksiGenap(angka int) bool {
	if angka % 2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	var masukan int

	fmt.Print("Masukkan angka: ")
	fmt.Scan(&masukan)

	fmt.Println("Apakah angka", masukan, "genap?")
	fmt.Println("Jawabannya adalah", deteksiGenap(masukan))
}