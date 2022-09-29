package main

import (
	"fmt"
	"time"
)

func main() {
	//make channel
	c1 := make(chan int)

	go func(c chan int) {
		for i := 1; i <= 5; i++ {
			fmt.Print("func goroutine %d starts sending data into the channel\n", i)
			c <- i
			fmt.Print("func goroutine %d after sending data into the channel\n", i)
		}
		close(c)
	}(c1)

	fmt.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	for v := range c1 {
		fmt.Println("main goroutine received value form channel:", v)
	}
}
