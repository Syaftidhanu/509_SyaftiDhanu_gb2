package main

import (
	"fmt"
)

func main() {
	var fruits1 = []string{"apple", "banana", "mango", "durian", "pineapple"}

	var fruits2 = fruits1[1:4]
	fmt.Printf("%#v\n", fruits2)

	var fruits3 = fruits1[0:]
	fmt.Printf("%#v\n", fruits3)

	var fruits4 = fruits1[:3]
	fmt.Printf("%#v\n", fruits4)

	var fruits5 = fruits1[:] //sama dengan fruit1[:len(fruits1)]
	fmt.Printf("%#v\n", fruits5)

}
