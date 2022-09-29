package main

import "fmt"

func main() {
	defer fmt.Println("defer function starts to excute")
	fmt.Println("Hai Everyone")
	fmt.Println("Welcome back to Go learning center")
}

//defer function dimana proses eksekusinya akan dikerjakan setelah function lain telah dieksekusi
//defer yang diletakan didalam func main
