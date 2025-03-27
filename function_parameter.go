package main


import "fmt"

// dalam golang, function juga bisa dijadikan sebagai parameter
// parameter function harus memiliki tipe data yang sama dengan function yang akan dijadikan sebagai parameter

// function di bawah ini memiliki 2 parameter yaitu tipe data string dan juga tipe data function dengan parameter tipe string
// function sebagai parameter dalam function ini memiliki parameter yang berupa nama (parameter dari function Jenis Air)
func JenisAir(nama string, filter func(string) string) {
	fmt.Println("Ini adalah air", filter(nama))
}


// function di bawah ini yang akan digunakan sebagai parameter dari function Jenis Air
func FilterAir(nama string) string {
	if nama == "Khamr" || nama == "Alkohol" {
		return "haram"
	} else {
		return nama
	}
}


// contoh lain dengan type declrataion untuk function
type Func func(int) string

func DeteksiPrima(angka int, filter Func) {
	fmt.Println("Angka", angka, filter(angka))
}

func isPrime(angka int) string {
	if angka < 2 {
		return "bukan angka prima"
	}

	for i := 2; i < angka; i++ {
		if angka % i == 0 {
			return "bukan angka prima"
		}
	}

	return "adalah angka prima"
}

func main() {
	JenisAir("Khamr", FilterAir)
	JenisAir("Susu", FilterAir)
	DeteksiPrima(2, isPrime)

	// memanggil function dalam bentuk value di variable
	deteksiPrima := isPrime
	DeteksiPrima(3, deteksiPrima)
}
