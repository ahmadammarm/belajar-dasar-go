package main

import "fmt"

// dalam golang terdapat operator perbandingan
// operator perbandingan adalah operator yang digunakan untuk membandingkan dua nilai
// operator perbandingan menghasilkan nilai boolean (true, false)
// operator perbandingan yang bisa digunakan adalah ==, !=, <, >, <=, >=
// untuk tipe data string yang menggunakan operasi lebih besar atau kecil maka akan dilihat dari urutan kamus
// untuk tipe data boolean, operator perbandingan hanya bisa menggunakan == dan !=
// operator perbandingan bisa digunakan untuk operasi antar variabel


func main() {
	var nilai1 string = "Ammar"
	var nilai2 string = "Adil"

	fmt.Println(nilai1 >= nilai2) // mengembalikan nilai true karena diurutkan secara kamus 
	fmt.Println(nilai1 <= nilai2) // mengembalikan nilai false karena diurutkan secara kamus
	fmt.Println(nilai1 == nilai2) // mengembalikan nilai false karena nilai1 dan nilai2 berbeda
	fmt.Println(nilai1 != nilai2) // mengembalikan nilai true karena nilai1 dan nilai2 berbeda
	fmt.Println(nilai1 > nilai2) // mengembalikan nilai true karena diurutkan secara kamus
	fmt.Println(nilai1 < nilai2) // mengembalikan nilai false karena diurutkan secara kamus

	var angka1 int = 10
	var angka2 int = 20

	fmt.Println(angka1 >= angka2) // mengembalikan nilai false karena angka1 lebih kecil dari angka2
	fmt.Println(angka1 <= angka2) // mengembalikan nilai true karena angka1 lebih kecil dari angka2
	fmt.Println(angka1 == angka2) // mengembalikan nilai false karena angka1 dan angka2 berbeda
	fmt.Println(angka1 != angka2) // mengembalikan nilai true karena angka1 dan angka2 berbeda
	fmt.Println(angka1 > angka2) // mengembalikan nilai false karena angka1 lebih kecil dari angka2
	fmt.Println(angka1 < angka2) // mengembalikan nilai true karena angka1 lebih kecil dari angka2

	var boolean1 bool = true
	var boolean2 bool = false

	fmt.Println(boolean1 == boolean2) // mengembalikan nilai false karena boolean1 dan boolean2 berbeda
	fmt.Println(boolean1 != boolean2) // mengembalikan nilai true karena boolean1 dan boolean2 berbeda
}