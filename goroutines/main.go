package main


// parallel programming adalah salah satu cara untuk meningkatkan performa program dengan memanfaatkan banyak CPU core secara bersamaan
// atau dengan bahasa ringannya adalah menjalankan beberapa proses secara bersamaan
// parallel programming berbeda dengan concurrent programming, pada parallel programming kita menjalankan beberapa proses secara bersamaan, sedangkan pada concurrent programming kita menjalankan beberapa proses secara bergantian
// contoh penerapan parallel adalah misalkan kita membuka beberapa tab di browser, kita bisa membuka beberapa tab secara bersamaan, sedangkan contoh penerapan concurrent adalah misalkan kita membuka beberapa tab di browser, kita bisa membuka beberapa tab secara bergantian

// pada parallel programming kita akan disinggungkan dengan 2 istilah yaitu process dan thread
// process adalah sebuah eksekusi program yang berjalan, sedangkan thread adalah sebuah eksekusi program yang berjalan di dalam process
// process terisolasi satu sama lain atau dengan kata lain, process tidak bisa saling berkomunikasi satu sama lain, sedangkan thread bisa saling berkomunikasi satu sama lain jika berada dalam satu process yang sama
// jadi bisa kita simpulkan bahwa thread adalah bagian dari process, dan process adalah bagian dari program

// contohnya, misalkan kita memiliki sebuah program yang bernama program1, program1 memiliki 2 process yaitu process1 dan process2, dan process1 memiliki 2 thread yaitu thread1 dan thread2, maka bisa kita simpulkan bahwa program1 memiliki 2 process dan 2 thread

// contoh kode
// func program1() {
// 	// process1
// 	go process1() // thread1}
// 	go process2() // thread2
// 	// process2
// 	go process3() // thread3
// 	go process4() // thread4
// }

// diatas adalah contoh kode yang menunjukkan bahwa kita memiliki 2 process yaitu process1 dan process2, dan process1 memiliki 2 thread yaitu thread1 dan thread2, dan process2 memiliki 2 thread yaitu thread3 dan thread4
// jadi bisa kita simpulkan bahwa program1 memiliki 2 process dan 4 thread

// dalam golang secara default menggunakan konsep concurrent programming, yaitu dengan menggunakan goroutine
// namun golang sekarang sudah bisa berjalan di cpu multi core, jadi kita bisa menggunakan parallel programming di golang

// cpu bound
// banyak penerapan algoritma dibuat yang hanya membutuhkan cpu saja, misalkan kita membuat algoritma untuk mencari bilangan prima, atau kita membuat algoritma untuk mencari fibonacci, atau kita membuat algoritma untuk mencari faktorial
// namun pada penerapan algoritma machine learning, kita membutuhkan banyak data untuk diolah, dan kita membutuhkan banyak memory untuk menyimpan data tersebut, jadi kita bisa menggunakan parallel programming untuk mengolah data tersebut
// oleh karena itu, machine learning biasanya lebih menggunakan gpu untuk mengolah data, karena gpu memiliki banyak core yang bisa digunakan untuk mengolah data secara bersamaan, sedangkan cpu hanya memiliki sedikit core yang bisa digunakan untuk mengolah data secara bersamaan

// i/o bound
// banyak penerapan algoritma yang membutuhkan i/o, misalkan kita membuat algoritma untuk mengolah data dari database, atau kita membuat algoritma untuk mengolah data dari file, atau kita membuat algoritma untuk mengolah data dari api
// walaupun terbantu dengan menggunakan parallel programming, namun akan lebih baik jika kita menggunakan concurrent programming
// karena pada concurrent programming kita bisa menggunakan channel untuk mengirim data dari satu goroutine ke goroutine yang lain, sedangkan pada parallel programming kita tidak bisa menggunakan channel untuk mengirim data dari satu goroutine ke goroutine yang lain
// jadi bisa kita simpulkan bahwa pada penerapan algoritma yang membutuhkan i/o lebih baik menggunakan concurrent programming, sedangkan pada penerapan algoritma yang membutuhkan cpu lebih baik menggunakan parallel programming

// goroutine dalam golang adalah sebuah thread ringan yang dikelola oleh Go Runtime
// goroutine itu sendiri berjalan dalam sebuah thread
// tidak seperti thread dalam parallel, goroutine berjalan secara concurrent

// sebenarnya, goroutine dijalankan oleh Go Scheduler dalam thread, yang dimana jumlah threadnya sebanyak GOMAXPROCS (jumlah core cpu yang digunakan)
// jadi goroutine tidak bisa dibilang sebagai pengganti thread, karena goroutine berjalan dalam thread
// namun yang mempermudah kita adalah, kita tidak perlu mengatur thread secara manual, karena Go Scheduler yang akan mengatur thread secara otomatis


func main() {
	
}