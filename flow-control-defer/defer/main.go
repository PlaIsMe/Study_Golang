package main

import "fmt"

func basicDefer() {
	// `defer` delays execution until the surrounding function returns
	defer fmt.Println("world") // This will run last

	fmt.Println("hello") // This runs first
}

func stackedDefer() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		// Deferred calls are pushed to a stack
		defer fmt.Println(i) // Arguments are evaluated immediately
	}

	fmt.Println("done") // This prints before the deferred calls
}

func main() {
	basicDefer()
	stackedDefer()
}

// Output:
// hello
// world
// counting
// done
// 9
// 8
// ...
// 0

// Notes:
// - Defer is useful for cleanup (e.g., closing files).
// - Calls execute in **last-in, first-out** (LIFO) order.
