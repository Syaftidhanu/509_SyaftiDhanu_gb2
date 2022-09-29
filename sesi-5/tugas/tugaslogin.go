package main

import (
	"errors"
	"fmt"

	"github.com/howeyc/gopass"
)

func validasiAkun(username string, password []byte) (string, error) {
	pl := string(password[:])
	if username != "syafti" || pl != "123" {
		return "", errors.New("Username/Password yang digunakan salah")
	}

	return "Berhasil Masuk", nil
}

func main() {
	var username string

	fmt.Println("Silahkan masukan username anda :")
	fmt.Scanln(&username)
	fmt.Println("Silahkan masukan password anda :")
	password, _ := gopass.GetPasswdMasked()

	if valid, err := validasiAkun(username, password); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(valid)
	}
}
