package main

import (
	"fmt"
	"strings"
)

func main() {
	var fruits1 = []string{"apple", "banana", "mango", "durian", "pineapple"}

	fmt.Println("fruits1 cap:", cap(fruits1)) //4
	fmt.Println("fruits2 len:", len(fruits1)) //4

	fmt.Println(strings.Repeat("#", 20))

	var fruits2 = fruits1[0:3]

	fmt.Println("fruits2 cap:", cap(fruits2)) //4
	fmt.Println("fruits2 len:", len(fruits2)) //3

	fmt.Println(strings.Repeat("#", 20))

	var fruits3 = fruits1[1:]

	fmt.Println("fruits3 cap:", cap(fruits3)) //3
	fmt.Println("fruits3 len:", len(fruits3)) //3

}
