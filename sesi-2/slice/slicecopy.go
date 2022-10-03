package main

import "fmt"

func main() {

	var fruits1 = []string{"apple", "banana", "mango"}

	var fruits2 = []string{"durian", "pineapple", "starfruit"}

	//argumen pertama dari fungsi copy merupakan slice destinasi
	//argumen kedua dari fungsi copy merupakan sumber
	//fungsi copy ini juga akan mengembalikan jumlah element yang berhasil dicopy
	nn := copy(fruits1, fruits2)

	fmt.Printf("Fruits1 =>", fruits1)
	fmt.Printf("Fruits2 =>", fruits2)
	fmt.Println("Copied elemeents =>", nn)
}
