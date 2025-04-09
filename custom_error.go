package main

import "fmt"

type errorValidation struct {
	Message string
}

func (validation *errorValidation) Error() string {
	return validation.Message
}

func SimpanData(data string) (string, error) {
	if data != "Ammar" {
		return "", &errorValidation{Message: "Data Bukan Ammar"}
	}

	return "Data Tersimpan", nil
}

func main() {
	errVa := &errorValidation{Message: "This is a custom error"}
	fmt.Println(errVa.Error()) // Output: This is a custom error
	fmt.Println(SimpanData("Ammar")) // Output: Data tidak boleh kosong
}