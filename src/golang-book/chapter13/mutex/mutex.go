package main

import (
	"fmt"
	"sync"
	"time"
)

// mutex: mutal exclusive lock
// locks a section of code to a single thread at a time
// is used to protect shared resources from non-atomic operations

func main() {
	m := new(sync.Mutex)

	for i := 0; i < 10; i++ {
		go func(i int) {
			m.Lock()
			fmt.Println(i, "start")
			time.Sleep(time.Second)
			fmt.Println(i, "end")
			m.Unlock()
		}(i)
	}

	var input string
	fmt.Scanln(&input)
}