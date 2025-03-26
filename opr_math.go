package main

import "fmt"

// operator matematika hanya berlaku untuk tipe data number (int, float)
// operator matematika yang bisa digunakan adalah +, -, *, /, %
// operator matematika juga bisa digunakan untuk operasi antar variabel
// operator matematika juga bisa digunakan untuk operasi antar variabel dan nilai langsung

func main() {
	var angka = 12 + 2
	var angka2 = 2	

	fmt.Printf("Hasil dari penjumlahan kedua variabel adalah %d", angka + angka2)

	// dalam golang juga terdapat augmented assignment
	// augmented assignment adalah operasi matematika yang dilakukan dengan menggunakan operator matematika dan assignment
	// augmented assignment bisa digunakan untuk operasi antar variabel

	angka += 2 // sama dengan angka = angka + 2
	fmt.Printf("\nHasil dari penjumlahan kedua variabel adalah %d", angka + angka2)

	// modulus adalah operator matematika yang digunakan untuk mendapatkan sisa pembagian

	var hasilbagi int8 = 10 % 2 // 10 dibagi 2 adalah 5 dengan sisa 0
	fmt.Printf("\nHasil dari modulus adalah %d", hasilbagi)

	// dalam golang juga terdapat unary operator
	// unary operator adalah operator matematika yang hanya membutuhkan satu operand
	// unary operator yang bisa digunakan adalah +, -, ++, --
	// unary operator bisa digunakan untuk operasi antar variabel

	var angka3 = 10
	angka3++ // sama dengan angka3 = angka3 + 1
	fmt.Printf("\nHasil dari unary operator adalah %d", angka3)
}