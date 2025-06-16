package main

import (
	"fmt"
	"math"
)

// Like for, Go's if does not need parentheses, but braces { } are required.

// Basic if
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i" // Recursive call for negative input
	}
	return fmt.Sprint(math.Sqrt(x))
}

// If with a short statement: 'v' is scoped to this if-else block
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v // v is only visible here
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim // v not visible here
}

func main() {
	// sqrt: handle real and imaginary roots
	fmt.Println(sqrt(2), sqrt(-4)) // Output: 1.4142135623730951 2i

	// pow: show value if under limit, or log otherwise
	fmt.Println(
		pow(3, 2, 10), // 9 < 10, returns 9
		pow(3, 3, 20), // 27 >= 20, prints message, returns 20
	)
}
