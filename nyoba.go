package main

import "fmt"

func main() {

	// var username string
	// var password string

	// fmt.Print("Username Anda: ", username)
	// fmt.Scan(&username)
	// fmt.Print("Kata Sandi Anda: ", password)
	// fmt.Scan(&password)

	// if username == "Coddy" {
	// 	if password == "1234" {
	// 		fmt.Println("Authentication successfully")
	// 	} else {
	// 		fmt.Println("Password incorrect")
	// 	}
	// } else {
	// 	fmt.Println("Username incorrect")
	// }

	// var buah string
	// var warna string

	// fmt.Print("Masukkan buah: ", buah)
	// fmt.Scan(&buah)

	// if buah == "Semangka" || buah == "semangka" {
	// 	if warna == "Merah" || warna == "merah" {
	// 		fmt.Print("Warna: ", warna)
	// 		fmt.Scan(&warna)
	// 		fmt.Println("Ini adalah semangka merah")
	// 	} else {
	// 		fmt.Println("Bukan semangka merah")
	// 	}
	// } else {
	// 	fmt.Println("Bukan buah semangka")
	// }

	// nested if
	// nested if adalah kondisi percabangan yang membutuhkan validasi lain meskipun kondisi pertama terpenuhi
	// jika kondisi pertama tidak terpenuhi maka tidak akan bisa masuk ke if berikutnya
	// var buah string
	// var warna string

	// fmt.Print("Masukkan nama buah: ", buah)
	// fmt.Scan(&buah)

	// if buah == "semangka" {

	// 	fmt.Print("Warna: ", warna)
	// 	fmt.Scan(&warna)

	// 	if warna == "merah" {
	// 		fmt.Println("Semangka merah")
	// 	} else {
	// 		fmt.Println("Bukan semangka merah")
	// 	}
	// } else {
	// 	fmt.Println("Bukan buah semangka")
	// }

	// // switch case
	// var monthlySalary int = 3000

	// switch monthlySalary * 12 {
	// 	case 36000:
	// 		fmt.Println("Junior")
	// 	case 48000:
	// 		fmt.Println("Senior")
	// 	default:
	// 		fmt.Println("default")
	// }

	// var nilai int
	// fmt.Scan(&nilai)

	// // operator bitwise
	// // operator bitwise adalah operator yang digunakan untuk melakukan operasi pada bit
	// // tanda << digunakan untuk melakukan operasi shift left
	// // tanda >> digunakan untuk melakukan operasi shift right

	// var nilai2 = nilai << 1
	// fmt.Println(nilai2)

	// var nilai3 int
	// fmt.Scan(&nilai3)

	// var nilai4 = ((nilai3 << 6) >> 1) + 18
	// fmt.Println(nilai4)

	// // asterisk loop
	// var a int
	// fmt.Scan(&a)

	// for index := 0; index < a; index++ {
	// 	for bintang := 0; bintang <= index; bintang++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	var namasaya string
	fmt.Print("Nama Anda: ")
	fmt.Scan(&namasaya)

	if namasaya == "Ammar" || namasaya == "ammar" {
		var jenisKelamin string
		fmt.Print("Jenis kelamin (L atau P): ")
		fmt.Scan(&jenisKelamin)
		if jenisKelamin == "L" {
			fmt.Println("Jenis kelamin benar")
		} else if jenisKelamin == "P" {
			fmt.Println("Jenis kelamin salah")
		} else {
			fmt.Println("masukan jenis kelamin tidak sesuai format")
		}
	} else {
		fmt.Println("Anda bukan Ammar")
	}


	slice := []int{1, 2, 3}

	fmt.Println(SequentialSearch(3, slice))

	var x int
	fmt.Scan(&x)

	for x < 10 {
		fmt.Println("Hello Coddy")
		x++
	}

	var a int = 4

	// print 4 baris angka 1-4 setiap barisnya sebanyak angka yang ada di baris tersebut
	for i := 1; i <= a; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(i)
		}
		fmt.Println()
	}

	array := [4]int{1, 2, 3, 4}

	for i := 0; i < len(array); i++ {
		fmt.Println(array[i] * 2)
	}
}

func SequentialSearch(yangdicari int, slice []int) bool {

	for index, elemen := range slice {
		if elemen == yangdicari {
			fmt.Println("Elemen", elemen, "ada di index", index)
			return true
		}
	}

	return false
}
