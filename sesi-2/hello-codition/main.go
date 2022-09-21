package main

import "fmt"

func main() {
	var currentYear = 2021

	if age := currentYear - 1999; age < 17 {
		fmt.Println("Kamu belum boleh membuat kartu sim")
	} else {
		fmt.Println("Kamu sudah boleh membuat sim")
	}
}
