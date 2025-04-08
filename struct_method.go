package main

import "fmt"

// struct adalah tipe data seperti tipe data lainnnya, dia bisa digunakan sebagai parameter dalam sebuah function
// jika kita ingin menambahkan method ke dalam struct, kita bisa juga
// method adalah function yang dimiliki oleh struct
// method dalam struct bisa diakses apabila struct tersebut sudah dideklarasikan dan ada datanya

type Product struct {
	Name string
	Price int
}

// method GetName adalah method yang dimiliki oleh struct Product untuk mengambil nama produk
func (product Product) GetName() string{
	return product.Name
}

// method GetPrice adalah method yang dimiliki oleh struct Product untuk mengambil harga produk
func (product Product) GetPrice() int{
	return product.Price
}

func main() {
	// mendeklarasikan struct Product
	product := Product{
		Name:  "Kaos",
		Price: 1000,
	}

	product2 := Product{
		Name:  "Kemeja",
		Price: 2000,
	}

	fmt.Println(product.GetName()) // memanggil method GetName dari struct Product
	fmt.Println(product2.GetName()) // memanggil method GetName dari struct Product
}