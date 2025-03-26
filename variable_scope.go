package main

import "fmt"

// varible scope adalah ruang lingkup dari sebuah variabel
// variabel yang dideklarasikan di luar fungsi bisa diakses oleh semua fungsi
// variabel yang dideklarasikan di dalam fungsi hanya bisa diakses oleh fungsi tersebut
// variabel yang dideklarasikan di dalam blok if, for, switch hanya bisa diakses oleh blok tersebut

// varible di bawah ini bisa diasup oleh semua fungsi
var namaku string = "Ammar"

// jika varible di luar fungsi sudah pernah dideklarasikan di file lain, maka akan terjadi error
// jika variable di luar fungsi dideklarasikan kembali di dalam fungsi, maka tidak akan terjadi error

func main() {
	var namaku string = "Ammar2"
	fmt.Println(namaku)
}