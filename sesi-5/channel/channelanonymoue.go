package main

import "fmt"

func main() {
	//make channel
	c := make(chan string)

	//array
	student := []string{"Airell", "Mailo", "Indah"}

	//function anonymous
	for _, v := range student {
		go func(student string) {
			fmt.Println("Student", student)
			result := fmt.Sprint("Hai, my name is %s", student)
			c <- result
		}(v)
	}

	for i := 1; i < 3; i++ {
		print(c)
	}

	close(c)
}

//fungsi menerima data melalui channel c
func print(c chan string) {
	fmt.Println(<-c)
}
