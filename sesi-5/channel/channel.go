package main

import "fmt"

func main() {
	//make channel
	c := make(chan string)

	//memanggil fungsi
	go introduce("airell", c)

	go introduce("Nanda", c)

	go introduce("Mailo", c)

	//menerima data dari channel
	msg1 := <-c
	fmt.Println(msg1)

	msg2 := <-c
	fmt.Println(msg2)

	msg3 := <-c
	fmt.Println(msg3)

}

//fungsi untuk mengirimkan data ke channel
func introduce(student string, c chan string) {
	result := fmt.Sprintf("Hai, my name is %s", student)

	c <- result
}
