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
// dalam golang, juga bisa untuk memberi nama untuk return value
// jika ingin mengembalikan nilai kembalian yang sama, maka bisa menggunakan tipe data yang sama


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
	
	// mengambil nilai kembalian dari function getHelloAndHi ke dalam variable
	// firstName, secondName := getHelloAndHi()
	// jika ingin mengambil salah satu nilai kembalian, maka bisa menggunakan _ (underscore) untuk mengabaikan nilai tersebut
	firstName , _ := getHelloAndHi()
	fmt.Println(firstName)

	// jika ingin mengembalikan semua nilai
	fmt.Println(getHelloAndHi())
	SequenctialSearch(5, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	fmt.Println(CekPalindrom("kasur rusak"))
}

func SequenctialSearch(angka int, slice []int) {
	for i := 0; i < len(slice); i++ {
		if angka == slice[i] {
			fmt.Println("Angka ditemukan di index ke-", i)
			break
		} else {
			fmt.Println("Angka tidak ditemukan")
		}
	}
}


// insertion sort adalah algoritma pengurutan sederhana yang membandingkan dua elemen pertama, kemudian membandingkan elemen ketiga dengan elemen kedua dan seterusnya
// insertion sort membandingkan elemen satu per satu dan memindahkan elemen tersebut ke posisi yang benar
// insertion sort cocok digunakan untuk jumlah elemen yang sedikit
// insertion sort memiliki kompleksitas waktu O(n^2)
// insertion sort memiliki kompleksitas ruang O(1)
func InsertionSort(slice []int) []int {
	for i := 1; i < len(slice); i++ {
		j := i - 1
		temp := slice[i]
		for j >= 0 && slice[j] > temp {
			slice[j+1] = slice[j]
			j--
		}
		slice[j+1] = temp
	}
	return slice
}

// function dengan named return value
func NamaFunction() (nama string) {
	nama = "Ammar"
	return nama
}

// palindrom adalah kata, frasa, angka, maupun susunan lainnya yang dapat dibaca dengan sama baik dari depan maupun belakang
// contoh palindrom adalah "kasur rusak", "ibu ratna antar ubi", "ada", "malam", "radar", "malas", "tanggat", "nurses run", "ibu membeli mobil", "ibu membeli balon"
func CekPalindrom(masukan string) bool {
	for i := 0; i < len(masukan); i++ {
		if masukan[0] != masukan[len(masukan) - 1] {
			return false
		} else if masukan[1] != masukan[len(masukan) - 2] {
			return false
		}
	}
	return true
}