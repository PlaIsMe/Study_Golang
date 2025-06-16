package main

import "fmt"

// Define a struct type
type Vertex struct {
	X int
	Y int
}

func main() {
	// Basic struct initialization
	v1 := Vertex{1, 2} // All fields
	v2 := Vertex{X: 1} // Partial fields, Y defaults to 0
	v3 := Vertex{}     // Zero value: X = 0, Y = 0
	p := &Vertex{1, 2} // Pointer to struct

	// Accessing fields
	fmt.Println(v1)   // {1 2}
	fmt.Println(v2.X) // 1
	fmt.Println(v3.Y) // 0

	// Modifying a field
	v1.X = 4
	fmt.Println(v1.X) // 4

	// Accessing fields through a pointer
	// Go automatically dereferences p.X to (*p).X
	fmt.Println(p.X) // 1
}
