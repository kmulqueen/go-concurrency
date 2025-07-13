package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)

}

func main() {
	var wg sync.WaitGroup

	words := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
	}

	wg.Add(len(words))

	for _, word := range words {
		go printSomething(word, &wg)
	}

	wg.Wait()

}
