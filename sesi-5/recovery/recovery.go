package main

import (
	"errors"
	"fmt"
)

func main() {
	defer catchErr()

	var password string

	fmt.Scanln(&password)

	if valid, err := validPassword(password); err != nil {
		//function panic
		panic(err.Error())
	} else {
		fmt.Println(valid)
	}
}

func catchErr() {
	if r := recover(); r != nil {
		fmt.Println("Error pccured:", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

func validPassword(password string) (string, error) {
	pl := len(password)
	if pl < 5 {
		return "", errors.New("password has to have more than 4 characters")
	}

	return "valid password", nil
}

//jika menggunakan panic semua proses goroutine akan terhenti
