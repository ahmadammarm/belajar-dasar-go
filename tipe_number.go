package main

import "fmt"

// tipe data number dalam golang ada 2 yaitu integer dan float
// integer adalah bilangan bulat
// float adalah bilangan desimal

// integer dalam golang terbagi menjadi 2 jenis yaitu signed dan unsigned
// signed adalah bilangan bulat yang bisa positif dan negatif
// unsigned adalah bilangan bulat yang hanya positif

// signed integer terbagi menjadi 4 jenis yaitu int8, int16, int32, int64
// unsigned integer terbagi menjadi 4 jenis yaitu uint8, uint16, uint32, uint64

// byte adalah alias dari uint8
// rune adalah alias dari int32
// int adalah alias dari minimal int32
// uint adalah alias dari minimal uint32

// int8 adalah bilangan bulat dengan panjang 8 bit (-128 sampai 127)
// int16 adalah bilangan bulat dengan panjang 16 bit (-32768 sampai 32767)
// int32 adalah bilangan bulat dengan panjang 32 bit (-2147483648 sampai 2147483647)
// int64 adalah bilangan bulat dengan panjang 64 bit (-9223372036854775808 sampai 9223372036854775807)

// uint8 adalah bilangan bulat dengan panjang 8 bit (0 sampai 255)
// uint16 adalah bilangan bulat dengan panjang 16 bit (0 sampai 65535)
// uint32 adalah bilangan bulat dengan panjang 32 bit (0 sampai 4294967295)
// uint64 adalah bilangan bulat dengan panjang 64 bit (0 sampai 18446744073709551615)

// penggunaan tipe data integer tergantung dari kebutuhan
// jika kita yakin nilai yang akan disimpan tidak akan melebihi batas dari int8, maka kita bisa menggunakan int8
// jika kita yakin nilai yang akan disimpan tidak akan melebihi batas dari int16, maka kita bisa menggunakan int16
// dan seterusnya
// karena jika kita secara default menggunakan int64 maka akan semakin banyak memori yang digunakan



// float32 adalah bilangan desimal dengan panjang 32 bit
// float64 adalah bilangan desimal dengan panjang 64 bit

// float32 memiliki presisi 7 digit
// float64 memiliki presisi 15 digit

func main() {
	var angka int8 = 127
	fmt.Printf("Angka %d\n", angka)

	var angkaDesimal float32 = 12.9
	fmt.Printf("Angka desimal %f\n", angkaDesimal)

	var angka2 int = 123
	fmt.Println("Angka 2 = ", angka2)

	var strings = "asda"
	fmt.Printf("String %s", strings)
	
	// fmt.Errorf("Error %s", strings)
}