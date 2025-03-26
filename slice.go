package main

import "fmt"

func main() {
	// dalam golang juga terdapat tipe data slice
	// slice adalah tipe data yang lebih fleksibel dari array
	// slice tidak memiliki panjang yang tetap
	// slice bisa juga bagian dari array maupun array itu sendiri
	// slice bisa diubah panjangnya
	// slice memiliki 3 properti, yaitu pointer, length, dan capacity
	// pointer adalah alamat memori dari elemen pertama slice
	// length adalah panjang dari slice
	// capacity adalah ditunjukkan dari elemen pertama slice hingga elemen terakhir array
	// slice juga bisa dibuat untuk mengubah data dalam array

	// cara membuat slice dari array
	var array = [5]int{1, 2, 3, 4, 5}
	var slice = array[1:4] // slice dari array index 1 sampai 4
	// dari slice tersebut, pointer terletak pada index 1 array, length adalah 3, dan capacity adalah 4
	fmt.Println(slice)

	var array2 = [5]int{1, 2, 3, 4, 5}
	var slice2 = array2[1:] // slice dari array index 1 sampai akhir
	// dari slice tersebut, pointer terletak pada index 1 array, length adalah 4, dan capacity adalah 4
	fmt.Println(slice2)

	var array3 = [5]int{1, 2, 3, 4, 5}
	var slice3 = array3[:4] // slice dari array index 0 sampai 4
	// dari slice tersebut, pointer terletak pada index 0 array, length adalah 4, dan capacity adalah 5
	fmt.Println(slice3)

	var array4 = [5]int{1, 2, 3, 4, 5}
	var slice4 = array4[:] // slice dari array index 0 sampai akhir
	// dari slice tersebut, pointer terletak pada index 0 array, length adalah 5, dan capacity adalah 5
	fmt.Println(slice4)

	// slice juga bisa dibuat tanpa array
	var slices = []int{1, 2, 3, 4, 5}
	fmt.Println(slices)


	// function dalam slice

	// len() digunakan untuk menghitung panjang slice
	fmt.Println(len(slices))

	// append() digunakan untuk menambahkan elemen pada slice
	// slice = append(slice, elemen yang ingin dimasukkan)
	slices = append(slices, 6)
	fmt.Println(slices)

	// cap() digunakan untuk menghitung kapasitas slice
	fmt.Println(cap(slices))
	
}