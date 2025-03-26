package main

import "fmt"

func main() {
	// constant adalah variabel yang nilainya tidak bisa diubah
	// constant dideklarasikan dengan menggunakan keyword const

	const nama string = "Ammar"
	fmt.Println(nama)

	// constant juga bisa dideklarasikan dengan multiple variable declaration
	const (
		namaDepan string = "Ammar"
		namaBelakang string = "Fathin"
	)

	const namaNama, alamat = "Ammar", "Indonesia"
	fmt.Println(namaNama, alamat)
	fmt.Println(namaDepan, namaBelakang)

	// dalam constant juga bisa menggunakan type inference
	const country = "Indonesia"

	// dalam constant jika tidak digunakan, maka tidak akan terjadi error
	// constant yang tidak digunakan tidak akan mempengaruhi performa dari program
	// constant yang tidak digunakan akan dihapus oleh compiler
}