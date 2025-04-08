package main

import "fmt"

type Alamat struct {
	Kota     string
	Provinsi string
	Negara   string
}

func GantiAlamat(alamat Alamat) string {
	negaraBaru := "Indonesia"
	alamat.Negara = negaraBaru
	return alamat.Negara
}

func GantiAlamatPointer(alamat *Alamat) string {
	negaraBaru := "Indonesia"
	alamat.Negara = negaraBaru
	return alamat.Negara
}

func main() {
	alamat := Alamat{
		"Kota Malang",
		"Jawa Timur",
		"",
	}

	// secara default struct alamat tersebut tidak berubah meskipun kita sudah memanggil fungsi GantiAlamat
	// karena kita hanya mengirimkan salinan dari alamat tersebut ke dalam fungsi GantiAlamat
	fmt.Println(GantiAlamat(alamat))

	// 

	fmt.Println(alamat)

	// mengganti dengan pointer
	// kita mengirimkan alamat dari struct tersebut ke dalam fungsi GantiAlamatPointer
	// sehingga kita bisa mengubah nilai dari struct tersebut
	
	fmt.Println(GantiAlamatPointer(&alamat))
	fmt.Println(alamat)
}