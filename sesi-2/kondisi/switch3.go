package main

import "fmt"

func main() {

	var score = 6

	switch {
	case score == 8:
		fmt.Println("perfect")
	case (score < 8) && (score > 3):
		fmt.Println("not bad")
		fallthrough
	case score < 5:
		fmt.Println("t is ok, but please study harder")
	default:
		{
			fmt.Println("Study Harder")
			fmt.Println("You don't have a good score yet")
		}
	}
}
