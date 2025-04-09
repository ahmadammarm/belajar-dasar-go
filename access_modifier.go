package main

import "fmt"

// dalam golang juga terdapat fitur access modifier, yaitu public dan private
// public adalah akses yang bisa diakses dari luar package
// public harus diinisiasi dengan huruf kapital
// private adalah akses yang hanya bisa diakses dari dalam package itu sendiri
// private harus diinisiasi dengan huruf kecil
// access modifier ini juga berlaku pada struct, function, dan variable

// contoh public dan private pada struct
// type Person struct { -> public
// type person struct { -> private

// contoh public dan private pada function
// func PublicFunction() { -> public
// func privateFunction() { -> private

// contoh public dan private pada variable
// var PublicVariable = 1 // public
// var privateVariable = 1 // private


func main() {
	fmt.Println("Access Modifier")
}