package main

import "fmt"

// defer function dalam go adalah function yang bisa kita tunda eksekusinya sampai function yang memanggilnya selesai dieksekusi
// defer function akan selalu dieksekusi walaupun terjadi error di function yang dieksekusi
// defer function juga bisa digunakan untuk menutup resource yang sudah dibuka sebelumnya

func loggin() {
	fmt.Println("Logging...")
}

func runApplication() {
	defer loggin() // defer function ini akan dieksekusi setelah function runApplication selesai dieksekusi
	// defer function ini akan dieksekusi walaupun terjadi error di function runApplication
	fmt.Println("Run Application")
}

// contoh defer dengan error

func runApplicationWithError() {
	defer loggin() // defer function ini akan dieksekusi setelah function runApplicationWithError selesai dieksekusi
	// defer function ini akan dieksekusi walaupun terjadi error di function runApplicationWithError
	fmt.Println("Run Application")
	panic("Error") // panic adalah error yang tidak bisa ditangani
	// jika panic terjadi, maka defer function tetap akan dieksekusi
}

func main() {
	runApplication()
	runApplicationWithError()
}