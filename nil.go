package main

import "fmt"

// nil adalah data kosong yang digunakan untuk menandai bahwa variable tersebut tidak memiliki nilai
// nil hanya bisa digunakan untuk tipe data pointer, slice, map, channel, interface, function, dan struct
// nil tidak bisa digunakan untuk tipe data int, string, float, bool, dan array

func NewMap(nama string) map[string]string {
	if nama == "" {
		return nil
	} else {
		return map[string]string{
			"Nama": nama,
		}
	}
}

func NewSlice (nama string) []string {
	if nama == "" {
		return nil
	} else {
		return []string{nama}
	}
}

func main() {
	// contoh kode nil

	var nilValue interface{}
	nilValue = nil
	fmt.Println(nilValue)

	fmt.Println(NewMap("Ammar"))
	// fmt.Println(NewMap(""))
	fmt.Println(NewSlice("Ammar"))

	slice := NewSlice("Ammar")
	slice = append(slice, "Ahmad")
	fmt.Println(slice)

}