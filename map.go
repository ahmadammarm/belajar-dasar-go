package main

import "fmt"

func main() {
	// dalam golang juga terdapat tipe data map
	// map adalah tipe data yang memiliki key dan value
	// map mirip seperti array, namun key pada map bisa ditentukan sendiri
	// key pada map harus unik, sedangkan value bisa duplikat
	// map bisa dibuat dengan cara map[key]value
	// map bisa dibuat dengan cara map[key]value{key:value}
	// map bisa dibuat dengan cara make(map[key]value)
	// map bisa dibuat dengan cara make(map[key]value, length)

	// cara membuat map
	var map1 map[string]string
	map1 = map[string]string{}
	map1["name"] = "Ahmad"
	map1["age"] = "20"
	fmt.Println(map1)

	person := map[string]string{
		"name": "Ahmad",
		"age":  "20",
	}
	fmt.Println(person)

	// function dalam map

	// len() digunakan untuk menghitung panjang map
	fmt.Println(len(person))

	// map[key] digunakan untuk mengakses value dari key
	fmt.Println(person["name"])

	// map[key] = value digunakan untuk mengubah value dari key
	person["name"] = "Ammar"
	fmt.Println(person)

	// delete(map, key) digunakan untuk menghapus key dari map
	delete(person, "name")

	fmt.Println(person)

	// make(map[key]value) digunakan untuk membuat map
	person2 := make(map[string]string)
	person2["name"] = "Ahmad"
	person2["age"] = "20"
	fmt.Println(person2)

	var masukan string
	fmt.Print("Masukkan nama mahasiswa: ")
	fmt.Scan(&masukan)

	mahasiswa1 := map[string]string {
		"nama": masukan,
		"nim": "nomor",
	}

	fmt.Println(mahasiswa1)

	var buku map[string]string
	buku = map[string]string{}
	buku["nama"] = "buku saku"
	buku["nomor"] = "12"

	delete(buku, "nomor")

	fmt.Println(buku)

}