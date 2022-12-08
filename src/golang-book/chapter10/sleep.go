package main

import (
	"time"
)

func sleep(n int64) {
	// time.After creates a channel tha will receive a value after the time given
	// the receiver will block for n seconds
	<-time.After(time.Second * time.Duration(n))
}

func main() {
	sleep(10)
}
