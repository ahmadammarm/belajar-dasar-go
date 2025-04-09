package main

import (
	"errors"
	"fmt"
)

// error merupakan built in interface yang digunakan untuk menangani error
// error merupakan interface yang memiliki satu method yaitu Error() string
// untuk menggunakan error kita bisa mengimport package errors
// kita bisa membuat error dengan menggunakan errors.New("pesan error")
func genap(angka int) (int, error) {
	if angka%2 == 0 {
		return angka, nil
	} else {
		return 0, errors.New("angka tidak genap")
	}
}

func main() {
	fmt.Println(genap(3))

	
	hasil, err := genap(3)

	// jika tidak ada error maka hasil akan berisi angka genap
	if err == nil {
		fmt.Println(hasil)
	// jika ada error maka err akan berisi pesan error yang diambil dari fungsi genap
	} else {
		fmt.Println(err.Error())
	}

}