package main

import (
	"fmt"
	"time"
)

// The channel will only send the message if the capacity is full

func pinger3(c chan<- int) {
	for i := 0; ; i++ {
		c <- 0
		c <- 1
		c <- 2
		c <- 3
		c <- 4
	}
}

func printer3(c <-chan int) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	// buffered channel of size 5
	c := make(chan int, 5)
	go pinger3(c)
	go printer3(c)

	var input string
	fmt.Scanln(&input)
}
