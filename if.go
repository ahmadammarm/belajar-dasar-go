package main

import "fmt"

// dalam golang juga mendukung if expression
// if expression adalah percabangan yang hanya mengeksekusi program tertentu jika kondisi terpenuhi
// dalam pengecekan if dimulai dari atas ke bawah


func main() {
	// if kondisi {hasil jika true} else {hasil jika false}
	// jika kondisi dalam if sudah terpenuhi maka program akan dieksekusi
	// jika ada else if, dibawah if yang true maka else if akan diabaikan
	var angka int = 10
	if angka == 10 {
		fmt.Println("Angka adalah 10")
	} else if angka == 10 {
		fmt.Println("Angka merupakan 10")
	} else {
		fmt.Println("Angka bukan 10")
	}

	// if expression juga bisa digunakan dalam variabel
	// variabel yang dideklarasikan dalam if expression hanya bisa digunakan dalam if expression
	// if ini disebut dengan if short statement
	if angka2 := 20; angka2 == 20 {
		fmt.Println("Angka adalah 20")
	} else {
		fmt.Println("Angka bukan 20")
	}

	var nama string
	fmt.Print("Nama Anda: ")
	fmt.Scan(&nama)

	if panjang := len(nama); panjang > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sesuai")
	}
	
}