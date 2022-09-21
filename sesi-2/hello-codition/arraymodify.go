package main

import "fmt"

func main() {

	var fruits = [3]string{"apel", "pisang", "mangga"}
	fruits[0] = "apple"
	fruits[1] = "pisang"
	fruits[2] = "mangga"

	fmt.Printf("%#v\n", fruits)

}
