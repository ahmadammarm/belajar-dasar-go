package main

import "fmt"

// variadic function
// parameter yang berada di posisi terakhir, memiliki kemampuan dijadikan sebuah varargs (variable arguments)
// varargs artinya kita bisa memasukkan lebih dari satu input ke dalam function tersebut seperti array
// perbedaan dari parameter biasa dengan tipe data array adalah pada tipe array, kita harus membuat array terlebih dahulu sebelum memasukkan array tersebut ke dalam function
// sedangkan varargs, kita bisa langsung memasukkan input berupa tipe data yang sama ke dalam function tersebut
// parameter dalam variadic function adalah slice


func average(angka ...int) int {
	
	total := 0

	for _, elemen := range angka {
		total += elemen
	}

	rataRata := total / len(angka)

	return rataRata
}

func main() {
	fmt.Println(average(2, 4, 6, 8, 10))
}