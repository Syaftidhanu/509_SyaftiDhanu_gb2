package main

import "fmt"

func main() {

	var score = 6

	switch {
	case score == 8:
		fmt.Println("perfect")
	case (score < 8) && (score > 3):
		fmt.Println("not bad")
	default:
		{
			fmt.Println("Study Harder")
			fmt.Println("You need to learn more")
		}
	}
}
