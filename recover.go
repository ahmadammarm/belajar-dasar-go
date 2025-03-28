package main

import "fmt"

func main() {
	// recover adalah function yang digunakan untuk menangani panic
	// recover hanya bisa digunakan di dalam defer function
	// jika panic terjadi, maka recover akan mengembalikan nilai dari panic tersebut
	// jika tidak ada panic, maka recover akan mengembalikan nil
	// jika recover dipanggil di luar defer function, maka recover tidak akan berfungsi
	// jika recover dipanggil di dalam defer function, maka recover akan mengembalikan nilai dari panic tersebut

	// mendefinisikan defer function dengan anonymous function
	// anonymous function ini akan dipanggil jika panic terjadi
	defer func() {
		message := recover()
		fmt.Println("Error:", message)
	}()

	panic("Error")
}