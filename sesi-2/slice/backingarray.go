package main

import (
	"fmt"
)

func main() {
	var fruits1 = []string{"apple", "banana", "mango", "durian", "pineapple"}

	var fruits2 = fruits1[2:4]

	fruits2[0] = "rambutan"

	fmt.Println("fruits1 => ", fruits1)
	fmt.Println("fruits2 => ", fruits2)

}
