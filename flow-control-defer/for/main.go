package main

import "fmt"

// Go has only one looping construct: the "for" loop.

// Basic for-loop: init; condition; post
func basicFor() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("basicFor:", sum)
}

// While-style loop: no semicolons
func whileStyle() {
	sum := 1
	// Init and post are optional: for ; sum < 1000;
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("whileStyle:", sum)
}

// Infinite loop: no condition
func infiniteLoop() {
	for {
		// Uncomment the next line to break and avoid endless loop
		// break
	}
}

func main() {
	basicFor()
	whileStyle()
	// infiniteLoop() // Don't run this unless you handle breaking
}
