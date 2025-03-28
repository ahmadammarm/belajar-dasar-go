package main

import "fmt"

type Mahasiswa struct {
	ID int
	Nama string
}

var mahasiswaList = []Mahasiswa{}

func AddMahasiswa(mahasiswa Mahasiswa) {
	for _, m := range mahasiswaList {
		if m.ID == mahasiswa.ID {
			fmt.Println("ID sudah ada")
			return
		}
	}

	if mahasiswa.ID < 1 {
		fmt.Println("ID tidak valid")
	} else {
		mahasiswaList = append(mahasiswaList, mahasiswa)
		fmt.Println("Berhasil menambahkan mahasiswa dengan ID", mahasiswa.ID, "dan nama", mahasiswa.Nama)
	}
}



func main() {
	
	var ID int
	fmt.Print("Masukkan ID Mahasiswa: ")
	fmt.Scan(&ID)

	var nama string
	fmt.Print("Masukkan Nama Mahasiswa: ")
	fmt.Scan(&nama)

	AddMahasiswa(Mahasiswa{ID: ID, Nama: nama})
	fmt.Println(mahasiswaList)

}