package main

import "fmt"

func main() {

	var fruits1 = []string{"apple", "banana", "mango"}

	var fruits2 = []string{"durian", "pineapple", "starfruit"}

	//fungsi untuk menggabungkan slice
	fruits1 = append(fruits1, fruits2...)

	fmt.Printf("%#v", fruits)
}
