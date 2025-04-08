package main

import "fmt"

// interface adalah tipe data abstrak yang tidak memiliki implementasi
// interface hanya memiliki method yang harus diimplementasikan oleh struct yang mengimplementasikan interface tersebut
// interface tidak memiliki data, hanya memiliki method
// interface digunakan untuk mendefinisikan contract atau kesepakatan antara struct dan interface yang diimplementasikan
// interface bisa digunakan untuk mendefinisikan method yang harus diimplementasikan oleh struct yang mengimplementasikan interface tersebut


type Animal interface {
	// method Speak adalah method yang harus diimplementasikan oleh struct yang mengimplementasikan interface Animal
	Speak() string

	// method dengan parameter
	speakWithName(name string) string

}

type Orang interface {
	AmbilNama() string
}

type orang struct {
	Nama string
	Umur int
}

func sayHellows(masukan Orang) {
	fmt.Println("Hello", masukan.AmbilNama())
}

func (orangs orang) AmbilNama() string {
	return orangs.Nama
}

type Person interface {
	GetName() string
}

type person struct {
	Name string
	Age  int
}

type animal struct {
	Name string
	Age  int
	Voice string
}

func (animal animal) speakWithName(name string) string {
	nama := animal.Name

	if nama != animal.Name {
		return "Nama tidak sesuai"
	} else {
		return fmt.Sprintf("Hewan %s berbunyi %s", nama, animal.Voice)
	}
	
}

func (animal animal) Speak() string {
	return animal.Voice
}

func sayHellow(nilai Person) string {
	return fmt.Sprintf("Hello %s", nilai.GetName())

}

func (persona person) GetName() string {
	return persona.Name
}

func main() {

	// mendeklarasikan struct animal
	animal := animal{
		Name:  "Kucing",
		Age:   2,
		Voice: "Meow",
	}

	fmt.Println(animal.Speak()) // memanggil method Speak dari struct animal
	fmt.Println(animal.speakWithName("Asada")) // memanggil method Speak dari struct animal dengan parameter name

	person := person {
		Name: "Ammar",
		Age:  20,
	}

	fmt.Println(sayHellow(person)) // memanggil method sayHellow dengan parameter person

	orang := orang{
		Nama: "Ammar",
		Umur: 21,
	}

	sayHellows(orang)

}