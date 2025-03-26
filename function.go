package main

import "fmt"

// function adalah blok kode yang berisi pernyataan yang akan dijalankan ketika function tersebut dipanggil
// function dapat digunakan untuk mengelompokkan kode agar lebih mudah dibaca dan dikelola
// function dapat menerima parameter dan mengembalikan nilai
// function dapat memiliki nama dan parameter, jika function tersebut memiliki nama maka function tersebut dapat dipanggil di dalam function lain
// parameter adalah variabel yang dideklarasikan di dalam tanda kurung setelah nama function
// parameter digunakan untuk menerima input dari luar function
// function dapat memiliki nilai kembalian, nilai kembalian digunakan untuk mengembalikan nilai dari function tersebut
// nilai kembalian dideklarasikan setelah tanda kurung kurawal dan sebelum tanda kurung kurawal penutup yang berupa tipe data yang akan dikembalikan
// jika function tersebut memiliki nilai kembalian maka function tersebut harus mengembalikan nilai tersebut
// jika function tersebut tidak memiliki nilai kembalian maka tidak perlu mengembalikan nilai
// function dapat memiliki banyak nilai kembalian
// function dapat memiliki nilai kembalian yang sama atau berbeda tipe data
// function yang ingin diimport ke dalam file lain harus memiliki nama function yang diawali dengan huruf kapital
// karena function tersebut merupakan public function
// function yang diawali dengan huruf kecil adalah private function, function tersebut hanya bisa diakses di dalam package tersebut


// contoh function tanpa parameter dan tanpa nilai kembalian
func sayHello() {
	fmt.Println("Hello")
}

// contoh function dengan parameter dan tanpa nilai kembalian
func sayHelloTo(name string) {
	fmt.Println("Hello", name)
}

// contoh function tanpa parameter dan dengan nilai kembalian
func getHello() string {
	return "Hello"
}

// contoh function dengan parameter dan nilai kembalian
func getHelloTo(name string) string {
	return "Hello " + name
}

// contoh function dengan banyak nilai kembalian
func getHelloAndHi() (string, string) {
	return "Hello", "Hi"
}

func main() {
	// pemanggilan function ke dalam function main
	sayHello()
	sayHelloTo("Golang")
	fmt.Println(getHello())
	fmt.Println(getHelloTo("Golang"))
	fmt.Println(getHelloAndHi())
}