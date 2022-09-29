package main

import "fmt"

func main() {
	//make channel
	c := make(chan string)

	//array
	student := []string{"Airell", "Mailo", "Indah"}

	for _, v := range student {
		go introduce(v, c)
	}

	for i := 1; i < 3; i++ {
		print(c)
	}

	close(c)
}

//fungsi menerima data melalui channel c
func print(c <-chan string) {
	fmt.Println(<-c)
}

//func introduce
func introduce(student string, c chan<- string) {
	result := fmt.Sprintf("Haim my name is %s", student)

	c <- result
}
