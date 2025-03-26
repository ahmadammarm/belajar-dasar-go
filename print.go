package main

import "fmt"

func main() {
	// dalam golang format print terdapat 3 jenis yaitu Print, Println, Printf
	// Print adalah function untuk menampilkan teks tanpa newline (enter)
	// Println adalah function untuk menampilkan teks dengan newline (enter)
	// Printf adalah function untuk menampilkan teks dengan format tertentu (mirip seperti pada bahasa C)
	// untuk menggunakan Printf kita harus menambahkan format specifier (%d, %s, %f, %t, %v)
	// format specifier adalah kode yang digunakan untuk menampilkan tipe data tertentu
	// %d adalah format specifier untuk integer
	// %s adalah format specifier untuk string
	// %f adalah format specifier untuk float
	// %t adalah format specifier untuk boolean
	// %v adalah format specifier untuk menampilkan semua tipe data

	// contoh penggunaan Print
	fmt.Print("Hello ")
	fmt.Print("Hello\n") // tanpa menggunakan newline (\n digunakan untuk newline)

	// contoh penggunaan Println
	fmt.Println("Hello")
	fmt.Println("Hello") // menggunakan newline

	// contoh penggunaan Printf
	fmt.Printf("%s %s", "Hello", "World\n") // menggunakan format specifier %s
	fmt.Printf("%s\n%s", "Hello", "World") // menggunakan format specifier %s dan newline

	// contoh penggunaan format specifier
	fmt.Printf("Angka %d\n", 10) // %d digunakan untuk integer
	fmt.Printf("Angka %f\n", 10.5) // %f digunakan untuk float
	fmt.Printf("Angka %s\n", "sepuluh") // %s digunakan untuk string
	fmt.Printf("Angka %t\n", true) // %t digunakan untuk boolean
	fmt.Printf("Angka %v\n", 10) // %v

	// untuk menggunakan newline dalam Printf kita bisa menambahkan setelah format specifier

	// kita bisa menggunakan format print sesuai dengan kebutuhan kita
}