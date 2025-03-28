package main

import "fmt"

// dalam go juga terdapat fitur panic yang digunakan untuk menghentikan eksekusi program
// panic adalah error yang tidak bisa ditangani, jika panic terjadi maka program akan berhenti
// saat panic function dipanggil, program akan terhenti, namun defer function yang ada di dalam program akan tetap dieksekusi

// contoh program panic dengan defer

func logging() {
	fmt.Println("Logging...")
}

func runApplications() {
	defer logging() // defer function ini akan dieksekusi setelah function runApplication selesai dieksekusi
	// defer function ini akan dieksekusi walaupun terjadi error di function runApplication
	fmt.Println("Run Application")
	// panic("Error") // panic adalah error yang tidak bisa ditangani
	// jika panic terjadi, maka defer function tetap akan dieksekusi
}

func main() {
	runApplications()
	
}
