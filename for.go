package main


import "fmt"

// dalam golang juga terdapat fitur perulangan
// perulangan adalah fitur yang digunakan untuk mengeksekusi kode program secara berulang kali
// perulangan for digunakan untuk melakukan perulangan dengan jumlah yang sudah ditentukan
// perulangan for juga bisa digunakan untuk melakukan perulangan tanpa batas

// dalam for, kita bisa menambahkan statement, yang dimana terdapat 2 statement yang bisa ditambahkan
// statement pertama adalah statement sebelum perulangan (init statement)
// statement kedua adalah statement setelah perulangan (post statement)

// contoh perulangan for dengan statement

// for index := 0; index < 5; index++ {
// 	fmt.Println(index)
// }

// dalam contoh di atas, ada init statement index := 0, post statement index++
// init statement digunakan untuk mendeklarasikan variabel index dengan nilai 0
// setelah itu dilakukan perulangan dengan kondisi index < 5
// setelah perulangan selesai, maka akan dieksekusi post statement index++

// alur dari eksekusi di atas adalah
// 1. init statement index := 0
// 2. cek kondisi index < 5
// 3. jika kondisi terpenuhi maka eksekusi kode program dengan mencetak nilai index (fmt.Println(index))
// 4. setelah eksekusi kode program selesai, maka eksekusi post statement index++
// 5. nilai index bertambah 1
// 6. kembali ke langkah 2
// 7. jika kondisi index < 5 tidak terpenuhi maka perulangan berhenti


func main() {

	var i int = 10

	for i > 0 {
		fmt.Println(i)
		i--
	}

	// for dalam for
	for i := 1; i <= 5; i++ {
		// dalam for kedua ini, ketika nilai i sudah bertambah, nilai j tetap dimulai dari 1
		// sehingga akan mencetak bintang sebanyak nilai i
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		// jika kondisi di j tidak terpenuhi maka akan keluar dari block for j dan mencetak baris baru
		fmt.Println()
	}

	for i := 0; i < 5; i++ {
		// dalam for ini, ketika nilai i bertambah, maka nilai j akan dimulai dari 0 dan mencetak nilai i
		// dalam for ini tidak mencetak nilai 0 karena nilai j dimulai dari 0
		// ketika kondisi j = 0 dan i = 0 dan bertemu dengan kondisi j < i maka tidak akan mencetak nilai 0
		// sehingga nilai 0 tidak akan dicetak
		for j := 0; j < i; j++ {
			fmt.Print(i)
		}
	}

	// perulangan sederhana
	counter := 1

	for counter <= 10 {
		fmt.Println(counter)
		counter++
	}

	// perulangan dengan for
	// for index := 0; index < 5; index++ {
	// 	fmt.Println(index)
	// }

	// perulangan dengan for dan break
	// kegunaan break adalah untuk menghentikan perulangan jika kondisi tertentu terpenuhi
	// for index := 0; index < 5; index++ {
	// 	if index == 3 {
	// 		break
	// 	}

	// 	fmt.Println(index)

	// for dengan continue
	// kegunaan continue adalah untuk menghentikan perulangan pada kondisi tertentu dan melanjutkan perulangan
	// sebagai contoh perulangan di bawah ini akan menghentikan perulangan ketika index = 3 dan melanjutkan perulangan hingga index 5
	// for index := 0; index < 5; index++ {
	// 	if index == 3 {
	// 		continue
	// 	}
	// 	fmt.Println(index)
	// }

	var masukan int
	fmt.Print("Masukkan keinginan Anda: ")
	fmt.Scan(&masukan)

	if masukan > 5 && masukan % 2 == 0 {
		for i := masukan; i > 5; i-- {
			for j := 0; j < i; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
		for i := 1; i <= 5; i++ {
			for j := 0; j < i; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
	} else if masukan % 2 != 0 {
		fmt.Println("Masukan adalah bilangan ganjil, tidak mencetak bintang.")
	} else {
		for i := 0; i < masukan; i++ {
			for j := 0; j <= i; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
	}

	// for juga bisa digunakan untuk data collection seperti array, slice, dan juga map
	
	// contoh penggunaan for untuk array
	// menghitung jumlah nilai dalam array
	var nilai = [5]int{90, 80, 85, 70, 95}
	var total int = 0
	for _, elemen := range nilai {
		total += elemen
	}

	fmt.Println("Total nilai adalah", total)

	// contoh penggunaan for untuk slice
	// menghitung rata-rata nilai dalam slice

	var slice = []int{90, 80, 85, 70, 95}
	var totalNilai int = 0

	for _, nilai := range slice {
		totalNilai += nilai
	}

	var rataRata float64 = float64(totalNilai) / float64(len(slice))
	fmt.Println("Rata-rata nilai adalah", rataRata)

	// contoh penggunaan for untuk map
	// menghitung jumlah nilai dalam map

	var nilaiMap = map[string]int{
		"Matematika": 90,
		"Fisika": 80,
		"Kimia": 85,
		"Biologi": 70,
	}

	var totalNilaiMap int = 0

	for _, nilai := range nilaiMap {
		totalNilaiMap += nilai
	}

	fmt.Println("Total nilai map adalah", totalNilaiMap)
	
	// untuk menghitung jumlah seperti implementasi diatas, tipe data dari elemen yang diiterasi harus numeric
	// jika tipe data elemen yang diiterasi bukan numeric, maka akan terjadi error

	// rumus untuk for range
	// for index, elemen := range collection {
	// 	// kode program
	// }

}