//defer yang diletakkan di luar fungsi main

package main

import "fmt"

func main() {
	callDeferFunc()
	fmt.Println("Hai Everyone!!")
}

func callDeferFunc() {
	defer deferFunc()
}

func deferFunc() {
	fmt.Println("Defer func starts to execute")
}
