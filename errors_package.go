package main

import (
	"errors"
	"fmt"
)

type CustomeError struct {
	Message string
}

func (errors *CustomeError) Error() string {
	return errors.Message
}

func deteksiuGenap(angka int) (bool, error) {
	if angka < 0 {
		return false, errors.New("angka tidak boleh negatif")
	} else if angka%2 == 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func main() {
	fmt.Println(deteksiuGenap(34))

	errors := &CustomeError{Message: "ini adalah error"}
	fmt.Println(errors.Error())
}
