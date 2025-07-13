package main

import "testing"

// ! This test should fail when running go test -race .
func Test_updateMessage(t *testing.T) {
	msg = "Hello world"

	wg.Add(2)
	go updateMessage("Hello from testing")
	go updateMessage("Hello again from testing")
	wg.Wait()

	if msg != "Hello from testing" {
		t.Error("Expected 'Hello from testing', but it was not there.")
	}
}
