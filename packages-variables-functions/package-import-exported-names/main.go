package main

// You can also write multiple import statements, like:
// import "fmt"
// import "math"

import (
	"fmt"
	"math"

	// By convention, the package name is the same as the last element of the import path.
	// For instance, the "math/rand" package comprises files that begin with the statement package rand.
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	// In Go, a name is exported if it begins with a capital letter
	// the math.pi is wrong, we should use math.Pi
	fmt.Println(math.Pi)
}
