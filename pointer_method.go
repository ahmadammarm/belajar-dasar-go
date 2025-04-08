package main

import "fmt"


// untuk mengubah nilai asli dari struct melalui method kita direkomendasikan untuk menggunakan pointer receiver
// karena jika kita menggunakan receiver biasa, method akan menerima salinan dari struct
// sehingga tidak akan mengubah nilai asli dari struct tersebut
// jika kita menggunakan salinan secara terus menerus, maka akan menghabiskan memori
// dan memperlambat kinerja program kita
type Orangs struct {
	Nama string
}

func (orang Orangs) Perkenalan() {
	orang.Nama = "Joko"
	fmt.Println("Halo, nama saya", orang.Nama)
}


func (orang *Orangs) PerkenalanPointer() {
	orang.Nama = "Joko"
	fmt.Println("Halo, nama saya", orang.Nama)
}

func main() {
	orang := Orangs{Nama: "Budi"}
	
	orang.Perkenalan()

	// secara default, method tidak mengubah nilai asli dari struct
	// karena method menerima salinan dari struct
	fmt.Println(orang.Nama)

	
	
	// ketikan ingin mengubah nilai asli dari struct, kita bisa menggunakan pointer receiver
	// dengan menambahkan tanda * pada receiver
	orang.PerkenalanPointer()
	fmt.Println(orang.Nama)
}