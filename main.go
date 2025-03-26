// package main menjadi entry point dari program go
// executable program harus memiliki package main dan func main
// func main adalah function yang dijalankan ketika program dijalankan
package main

import "fmt"


// dalam golang untuk menjalankan program harus dicompile menjadi binary / executable file terlebih dahulu
// dengan cara mengetikkan perintah go build pada terminal (go build main.go)
// perintah go build hanya mengompile function main saja
// setelah itu akan muncul file binary / executable file dengan nama main
// untuk menjalankan file binary / executable file tersebut, kita bisa mengetikkan perintah ./main

// atau kita bisa langsung menjalankan program dengan periintah go run main.go
// perintah go run main.go akan langsung mengompile dan menjalan program tanpa harus membuat file binary / executable file
// cara ini hanya cocok ketika dalam tahap development
// jika sudah dalam tahap production, lebih baik menggunakan perintah go build terlebih dahulu
// karena file binary / executable file akan lebih cepat dalam menjalankan program
func main() {
	fmt.Println("Program apa yang biasa kamu tuliskan pertama kali ketika belajar bahasa pemrograman baru?")
	fmt.Println("Hello you")
	fmt.Println("Kenapa kok tidak Hello world seperti pada umumnya?")
	fmt.Println("Because you are my world")
}