package main

import "fmt"

func main() {

	var fruits = make([]string, 3)

	//fungsi append merupakan salah satu cara untuk untuk menambahkan element pada array
	//fruits biru nama variable
	//variabel pertama pada fungsi append adlh slice yang ingin ditambahkan
	fruits = append(fruits, "apple", "banana", "mango")

	fmt.Printf("%#v", fruits)
}
