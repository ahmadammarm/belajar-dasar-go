package main

import "fmt"

// dalam golang juga terdapat closure
// closure adalah function yang memiliki akses ke variable di luar function tersebut namun masih dalam scope yang sama


func main() {
	counter := 0
	// mendefinisikan vatriable increment yang berisi anonymous function (function tanpa nama)
	// anonymous function ini memiliki akses ke variable counter di luar function tersebut
	// anonymous function ini bisa mengakses variable counter karena variable counter berada di luar function tersebut
	increment := func() {
		counter++
		fmt.Println(counter)
	}

	increment()
	increment()
	increment()
}