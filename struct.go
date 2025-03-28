package main

import "fmt"


// struct dalam go adalah tipe data yang digunakan untuk menyimpan data yang memiliki beberapa field
// struct mirip dengan class di bahasa pemrograman lain
// struct tidak memiliki method, namun kita bisa menambahkan method ke dalam struct dengan cara mendefinisikan function di luar struct
// struct juga bisa memiliki anonymous field, yaitu field yang tidak memiliki nama
// secara default, struct tidak bisa langsung digunakan
// untuk menambah data ke dalam struct, kita bisa memasukkan struct tersebut ke dalam value sebuah variable
// jika kita langsung mencetak sebuah struct tanpa data, maka secara default akan terisi dengan default value setiap tipe data
// default value untuk string adalah "", untuk int adalah 0, untuk float adalah 0.0, untuk bool adalah false

type Mahasiswas struct {
	Nama, NIM string
	IPK        float64
}


func main() {
	// memasukkan data ke struct dengan cara mendeklarasikan struct tersebut ke dalam variable
	var mahasiswa Mahasiswas
	mahasiswa.Nama = "Ammar"
	mahasiswa.NIM = "220535601431"
	mahasiswa.IPK = 3.9

	fmt.Println(mahasiswa)

	// atau bisa dengan cara mendeklarasikan struct tersebut ke dalam variable dengan data yang sudah ada
	mahasiswa2 := Mahasiswas{
		Nama: "Ahmad",
		NIM:  "220535601432",
		IPK:  3.8,
	}

	fmt.Println(mahasiswa2)


}