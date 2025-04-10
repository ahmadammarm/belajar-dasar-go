package main

// os package adalah package yang digunakan untuk berinteraksi dengan sistem operasi
// seperti membaca dan menulis file, mengatur environment variable, dan lain-lain.
// Pada contoh ini kita akan menggunakan os.Args untuk mendapatkan argumen yang diberikan saat menjalankan program
// dan os.Hostname untuk mendapatkan hostname dari sistem operasi.


import (
	"errors"
	"fmt"
	"os"
)

func main() {
	oss := os.Args
	
	for _, elemen := range oss {
		fmt.Println(elemen)
	}

	hostname, err := os.Hostname()

	if err == nil {
		fmt.Println(hostname)
	} else {
		errors.New("error getting hostname")
	}
	
}