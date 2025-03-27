package main

import "fmt"


// function dalam golang juga dapat dijadikan sebagai value dalam variabel
// dalam value sebuah variabel, bisa diisi dengan nama functionnya tanpa tanda kurung
func CekPalindroms(masukan string) bool {
	for i := 0; i < len(masukan); i++ {
		if masukan[0] != masukan[len(masukan)-1] {
			return false
		} else if masukan [1] != masukan[len(masukan)-2] {
			return false
		}
	}
	return true
}


func main() {
	// pemanggilan function ke dalam value variabel
	// variabel cekPalindrom diisi dengan nama function CekPalindroms
	// jika kita menggunakan tanda kurung, maka kita berarti memanggil function tersebut dan bukan mengisi value variabel
	cekPalindrom := CekPalindroms
	fmt.Println(cekPalindrom("kasur rusak"))
}