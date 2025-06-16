package main

import "fmt"

// Zero values:
// - 0 for numbers
// - false for bool
// - "" for strings
var c, python, java bool
var i, j int = 1, 2

func main() {
	k := 3    // Short variable declaration (only inside functions)
	var i int // Shadows global 'i', defaults to 0

	// Type conversion
	ii := 42          // int
	ff := float64(ii) // converted to float64
	uu := uint(ff)    // converted to uint

	// Type inference
	x := 42           // int
	y := 3.14         // float64
	z := 0.867 + 0.5i // complex128

	fmt.Println(i, c, python, java, k)
	fmt.Println(ii, ff, uu)
	fmt.Printf("%T %T %T\n", x, y, z) // Print inferred types
}
