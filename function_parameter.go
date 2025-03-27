package main


import "fmt"

// dalam golang, function juga bisa dijadikan sebagai parameter
// parameter function harus memiliki tipe data yang sama dengan function yang akan dijadikan sebagai parameter

// function sayHelloWithFilter memiliki 2 parameter, yaitu name bertipe string dan filter bertipe function yang menerima parameter string dan mengembalikan string
// function sayHelloWithFilter akan mencetak "Hello" dan nama yang sudah difilter oleh function filter
func sayHelloWithFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

// function spamFilter memiliki 1 parameter, yaitu name bertipe string
// function spamFilter akan memfilter nama yang mengandung kata "Anjing" menjadi "...."
func spamFilter(name string) string {
	if name == "Anjing" {
		return "...."
	} else {
		return name
	}
}



func main() {
	sayHelloWithFilter("Anjing", spamFilter)
	sayHelloWithFilter("Budi", spamFilter)
}
