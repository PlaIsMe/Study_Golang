package main

import (
	"fmt"
	"time"
)

// MyError is a custom error type that includes time and message
type MyError struct {
	When time.Time
	What string
}

// Error implements the error interface for MyError
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

// run returns a custom error
func run() error {
	return &MyError{
		When: time.Now(),
		What: "something went wrong",
	}
}

func main() {
	if err := run(); err != nil {
		// Since MyError implements Error(), this prints the custom message
		fmt.Println("Error occurred:", err)
	}
}
