package main

import "fmt"

func main() {

	//multiple declaration
	var name, age, address = "Syafti", 23, "Bekasi"
	first, second, third := "1", 2, "3"

	fmt.Println(name, age, address)
	fmt.Println(first, second, third)

	var firstVariable string

	_, _, _, _ = firstVariable, name, age, address

	my_first, my_second := 1, "2"
	fmt.Printf("Tipe data variabel first adalah %T \n", my_first)
	fmt.Printf("Tipe data variabel second adalah %T \n", my_second)

	//menggunakan verb
	fmt.Printf("Hallo nama ku %s, umurku adalah %d, dan aku tinggal di %s", name, age, address)
}
