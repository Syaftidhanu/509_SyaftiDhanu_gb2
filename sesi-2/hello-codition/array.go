package main

import "fmt"

func main() {

	var numbers [4]int
	numbers = [4]int{1, 2, 3, 4}

	var string = [3]string{"Airell", "Nanda", "Mailo"}

	fmt.Printf("%#v\n", numbers)
	fmt.Printf("%#v\n", string)
}
