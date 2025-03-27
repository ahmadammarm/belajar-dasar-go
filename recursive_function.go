package main

import "fmt"

// recursive function dalam go adalah function yang memanggil dirinya sendiri
// recursive function biasanya digunakan untuk menyelesaikan permasalahan yang bisa dipecahkan dengan cara rekursif
// contoh paling sederhana dari recursive function adalah function faktorial (5! = 5 * 4 * 3 * 2 * 1)

func faktorial(angka int) int {
	if angka == 1 {
		return 1
	} else if angka == 0 {
		return 0
	} else {
		return angka * faktorial(angka - 1)
	}
} 

// selain faktorial, recursive function juga bisa digunakan untuk menyelesaikan permasalahan seperti fibonacci, permutasi, dll

// fibonacci adalah deret angka dimana angka selanjutnya adalah hasil penjumlahan dari 2 angka sebelumnya

func fibonacci(n int) int {
	if n <= 1 {
		return n
	} else {
		return fibonacci(n - 1) + fibonacci(n - 2)
	}
}

// permutasi adalah cara mengatur objek yang berbeda menjadi urutan tertentu

func permutasi(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * permutasi(n - 1)
	}
}

func main() {
	fmt.Println(faktorial(4))
	fmt.Println(fibonacci(5))
	fmt.Println(permutasi(5))
}