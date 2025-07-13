package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()
	m.Lock()
	msg = s
	m.Unlock()
}

// * Not sure what the result will be because we still aren't waiting
// * for 1 GoRoutine to finish before the other (it may be cosmos or universe)
// * but the important thing is we're accessing data safely (thread-safe operation)
// * No race condition
func main() {
	msg = "Hello World"

	var mutex sync.Mutex

	wg.Add(2)
	go updateMessage("Hello universe", &mutex)
	go updateMessage("Hello cosmos", &mutex)
	wg.Wait()

	fmt.Println(msg)
}
