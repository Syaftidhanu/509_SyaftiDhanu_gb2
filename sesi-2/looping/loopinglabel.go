package main

import "fmt"

func main() {
outerloop:
	for i := 0; i < 3; i++ {
		fmt.Println("Perulangan ke - ", i+1)
		for j := 0; j < 3; j++ {
			if i == 2 {
				break outerloop
			}
			fmt.Print(j, " ")
		}
		fmt.Println("\n")
	}
}
