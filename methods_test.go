package main

import (
	"fmt"
	"testing"
)

// Run test document `go test -run=AddTwoNumbers -bench=` Using regular expressions to run tests relating back to the add two numbers metho
func TestAddTwoNumbers(test *testing.T) { // B struct provides benchmarking methods
	fmt.Println("Began testing add two numbers method")
	expected := 4
	result := AddTwoNumbers(2, 2)

	if expected != result {
		test.Error("Result did not equal the expected output") // Allows you to provide error logs
	}
}
