package main

import "fmt"

// dalam golang percabangan selain if ada switch expression
// switch expression adalah percabangan yang mengecek kondisi yang sama
// switch expression digunakan jika terdapat banyak kondisi yang sama
// switch expression juga bisa digunakan dalam variabel
// sama seperti if, jika kondisi sudah terpenuhi maka program akan dieksekusi dan tidak akan mengecek kondisi lainnya
// jika tidak ada kondisi yang terpenuhi maka akan dieksekusi default sama seperti else
// jika tidak ada kondisi yang terpenuhi dan tidak ada default maka program tidak akan dieksekusi

func main() {
	var nama string = "Ammar"

	switch nama {
	case "Ammar":
		fmt.Println("Nama adalah Ammar")
	case "Fathin":
		fmt.Println("Nama adalah Fathin")
	default:
		fmt.Println("Nama tidak ditemukan")
	}

	var angka int = 100

	switch {
	case angka == 100:
		fmt.Println("Angka adalah 100")
	case angka == 200:
		fmt.Println("Angka adalah 200")
	default:
		fmt.Println("Angka tidak ditemukan")
	}

	// switch juga bisa dideklarasikan dengan short statement
	// variabel yang dideklarasikan dalam switch expression hanya bisa digunakan dalam switch expression
	switch angka2 := 100; angka2 {
	case 100:
		fmt.Println("Angka adalah 100")
	case 200:
		fmt.Println("Angka adalah 200")
	default:
		fmt.Println("Angka tidak ditemukan")
	}

	


}
