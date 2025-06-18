package main

import "fmt"

// Index is a generic function that returns the index of value x in slice s.
// If x is not found in s, it returns -1.
//
// T is a type parameter constrained by the built-in 'comparable' constraint,
// which means T must support == and != operations (e.g. int, string, float64).
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are of type T, which is comparable
		if v == x {
			return i // return the index if a match is found
		}
	}
	return -1 // return -1 if x is not found in s
}

func main() {
	// Example 1: Using Index with a slice of integers
	si := []int{10, 20, 15, -10}
	fmt.Println("Index of 15 in si:", Index(si, 15)) // Output: 2

	// Example 2: Using Index with a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(`Index of "hello" in ss:`, Index(ss, "hello")) // Output: -1 (not found)

	// Example 3: Using Index with a slice of float64
	sf := []float64{1.1, 2.2, 3.3}
	fmt.Println("Index of 2.2 in sf:", Index(sf, 2.2)) // Output: 1

	// Example 4: Using Index with a slice of bools
	sb := []bool{true, false, true}
	fmt.Println("Index of false in sb:", Index(sb, false)) // Output: 1
}
