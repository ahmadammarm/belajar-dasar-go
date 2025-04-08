package main

import "fmt"

// interface kosong adalah interface yang tidak memiliki method apapun
// hal ini dapat membuat interface kosong menjadi tipe data yang dapat digunakan untuk menyimpan nilai apapun
// interface kosong dapat digunakan untuk menyimpan nilai apapun, baik itu struct, int, string, float, dll
// interface kosong juga disebut dengan any
// interface kosong bisa diberi type assertion untuk mengubah tipe data yang ada di dalam interface kosong tersebut

type EmptyInterface interface{}

// ada beberapa penggunaan interface koosng dalam go
// 1. fmt.Println(interface{}) adalah function yang digunakan untuk mencetak nilai apapun ke layar
// 2. panic(interface{}) adalah function yang digunakan untuk menghentikan program dan mencetak error ke layar
// 3. recover() interface{} adalah function yang digunakan untuk menangkap panic dan mengembalikan nilai ke caller

func main() {
	// mendeklarasikan interface kosong
	var emptyInterface EmptyInterface

	// memasukkan nilai ke dalam interface kosong
	emptyInterface = "Hello World"
	fmt.Println(emptyInterface)

	emptyInterface = 123
	fmt.Println(emptyInterface)

	emptyInterface = 3.14
	fmt.Println(emptyInterface)

	emptyInterface = true
	fmt.Println(emptyInterface)

	emptyInterface = []string{"A", "B", "C"}
	fmt.Println(emptyInterface)

	emptyInterface = map[string]int{"A": 1, "B": 2, "C": 3}
	fmt.Println(emptyInterface)

	// type assertion
	// type assertion adalah cara untuk mengubah tipe data yang ada di dalam interface kosong

	// contoh kode
	

}