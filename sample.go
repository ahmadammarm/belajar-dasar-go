package main

import "fmt"

// dalam golang, multiple main function tidak diperbolehkan
// jika ada multiple main function, maka akan terjadi error pada saat compile (go build)
// karena tidak tahu main function mana yang akan dijalankan
// namun jika ingin menjalankan kode tersebut kita bisa menjalankan di setiap filenya (baik go run <nama file> maupun go build <nama file>)

func main() {
	fmt.Println("Oke bossku")
}