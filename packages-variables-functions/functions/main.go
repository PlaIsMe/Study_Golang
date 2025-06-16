package main

import (
	"fmt"
	"math"
)

// Basic function: two int inputs, one int output
func add(x, y int) int {
	return x + y
}

// Returns two strings, in reverse order
func swap(x, y string) (string, string) {
	return y, x
}

// Named return values: x and y are defined in signature
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // returns x, y implicitly
}

func main() {
	fmt.Println(add(42, 13))            // => 55
	fmt.Println(swap("hello", "world")) // => world hello
	fmt.Println(split(17))              // => 7 10

	// -------- Function Values --------
	// Functions are values — assign to variable and call later
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(3, 4)) // => 5

	// Function as parameter
	fmt.Println(compute(hypot))    // hypot(3,4) = 5
	fmt.Println(compute(math.Pow)) // pow(3,4) = 81

	// -------- Closures --------
	// adder returns a function that "remembers" its own sum
	pos, neg := adder(), adder()
	for i := 0; i < 5; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}

// compute takes a function (two float64 → one float64)
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// adder returns a closure that captures "sum" variable
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x // remembers and modifies outer sum
		return sum
	}
}
