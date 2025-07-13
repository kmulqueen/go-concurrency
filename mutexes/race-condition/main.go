package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

// ! This has a race condition. Run with `go run -race` to see warning.
func main() {
	msg = "Hello World"

	wg.Add(2)
	go updateMessage("Hello universe")
	go updateMessage("Hello cosmos")
	wg.Wait()

	fmt.Println(msg)
}
