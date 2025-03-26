package main

import "fmt"

func main() {
	// variabel adalah tempat untuk menyimpan data
	// variabel digunakan agar kita bisa mengakses data tersebut dengan mudah
	// dalam golang, variabel hanya bisa menyimpan tipe data yang sama
	// jika ingin menyimpan tipe data yang berbeda, kita harus membuat variabel baru
	// untuk membuat variabel kita bisa menggunakan keyword var
	// var nama_variabel tipe_data = nilai

	var nama string = "Ammar"

	var name string
	name = "Ammar"

	fmt.Println(name)
	fmt.Println(nama)

	// kita juga bisa membuat variabel tanpa menuliskan tipe datanya
	// golang akan otomatis menentukan tipe data dari nilai yang diberikan

	var age = 20
	fmt.Println(age)

	// kita juga bisa menuliskan variable dengan type inference
	// type inference adalah kemampuan dari golang untuk menentukan tipe data dari nilai yang diberikan

	country := "Indonesia"

	fmt.Println(country)

	var namaSaya string

	namaSaya = "Ammar"
	fmt.Println(namaSaya)

	// kita juga bisa menggunakan variabel yang sudah dideklarasikan sebelumnya sebagai nilai dari variabel baru
	// namun tidak mengubah nilai dari variabel yang sudah dideklarasikan sebelumnya
	// ini disebut dengan pass by value
	var namaSaya2 string = namaSaya
	namaSaya2 = "Ammar2"
	fmt.Println(namaSaya2)

	// multiple variable declaration
	// kita bisa mendeklarasikan beberapa variabel sekaligus
	var (
		namaDepan string = "Ammar"
		namaBelakang string = "Fathin"
	)

	var namaNama, alamat = "Ammar", "Indonesia"
	fmt.Println(namaNama, alamat)

	fmt.Println(namaDepan, namaBelakang)


}