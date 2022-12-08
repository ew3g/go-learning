package main

import (
	"fmt"
	"time"
)

// Channel is a way for two goroutines to communicate with one another
// and synchronize their execution
// We can also specify a direction on a channel
// restricting it to either sending or receiving

// pinger2 can only send messages to the channel
func pinger2(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

// ponger2 can only send messages to the channel
func ponger2(c chan<- string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

// printer2 can only receive messages from the channel
func printer2(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	c := make(chan string)
	go pinger2(c)
	go ponger2(c)
	go printer2(c)

	var input string
	fmt.Scanln(&input)
}
