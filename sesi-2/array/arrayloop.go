package main

import (
	"fmt"
)

func main() {

	var fruits = [3]string{"apple", "banana", "mango"}
	//cara pertama
	for i, v := range fruits {
		fmt.Printf("Index : %d, value: %s\n", i, v)
	}
}
