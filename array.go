package main

import "fmt"

// dalam golang juga terdapat tipe data array
// array adalah kumpulan data yang memiliki tipe data yang sama
// array memiliki index yang dimulai dari 0
// array memiliki panjang yang tetap dan tidak bisa diubah
// panjang array harus ditentukan saat deklarasi
// data dalam array tidak dapat dihapus, hanya bisa diubah

func main() {
	// pada saat mendeklarasikan array dengan tipe data yang diinginkan, maka dalam tipe data tersebut harus juga didefinisikan jumlah elemen yang akan disimpan
	var array [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)

	var array2 [3]string = [3]string{"Ahmad", "Ammar", "Musyaffa"}
	fmt.Println(array2)
	fmt.Println(array2[0])

	// atau bisa juga dideklarasikan seperti ini
	var array3 [3]int
	array3[0] = 1
	array3[1] = 2
	array3[2] = 3

	fmt.Println(array3)

	array3[2] = 12

	fmt.Println(array3)

	// namun best practice dalam mendeklarasikan array adalah dengan cara seperti ini
	array4 := [3]int{1, 2, 3}
	fmt.Println(array4)

	// jika kita mendeklarasikan array namun elemennya kurang dari jumlah elemen yang dideklarasikan, maka elemen yang tidak diisi akan bernilai default dari tipe data yang dideklarasikan
	// jika integer maka defaultnya 0
	// jika string maka defaultnya ""
	// jika boolean maka defaultnya false
	// jika float maka defaultnya 0.0

	array5 := [2]bool{true}
	fmt.Println(array5)

	// function dalam array
	// len() digunakan untuk menghitung panjang array
	fmt.Println(len(array4))

	// array[index] digunakan untuk mengakses elemen array
	fmt.Println(array4[0])

	// array[index] = value digunakan untuk mengubah nilai elemen array
	array4[0] = 10
	fmt.Println(array4)

}