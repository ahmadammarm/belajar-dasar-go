package main

import "fmt"

func main() {
	// tipe data string adalah tipe data yang digunakan untuk menyimpan teks
	// string dalam golang adalah kumpulan dari karakter yang diapit oleh tanda kutip ("")
	fmt.Println("This is string")


	// dalam string juga memiliki function yang bisa digunakan untuk manipulasi string
	// function tersebut adalah len, untuk mengetahui panjang dari string
	fmt.Println(len("Lets f*cking go"))


	// ada yang lain yaitu "string"[index] untuk mengambil karakter pada index tertentu
	// index tersebut dimulai dari 0
	// pada pengambilan index dari string akan mengembalikan byte dari huruf di index tersebut
	fmt.Println("Hello World"[0])
	fmt.Println("Hello World"[1:3]) // mengambil karakter dari index 1 sampai 3 (index 3 tidak diambil)

	fmt.Println(len("Hello World"[0:5])) // mengambil panjang dari karakter yang diambil dari index 0 sampai 5 (index 5 tidak diambil)
}