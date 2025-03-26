package main

import "fmt"

// dalam golang kita bisa membuat alias dari tipe data yang sudah ada
// alias ini bisa digunakan untuk membuat tipe data baru yang lebih mudah dibaca
type nama string

func main() {
	var namaSaya nama = "Ammar"
	fmt.Println(namaSaya)
}