package main

import (
	"fmt"
	"math"
)

// Vertex defines a simple 2D point
type Vertex struct {
	X, Y float64
}

// --- Methods ---

// Scale modifies the receiver's values by a factor (needs pointer receiver to apply changes)
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Abs calculates the distance from the origin (does not need to modify the receiver)
// But we use a pointer receiver here for consistency and to avoid copying large structs
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// --- Functions ---

// Function version of Scale (must take a pointer explicitly)
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Function version of Abs (must take a value explicitly)
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	// ----- Value -----
	v := Vertex{3, 4}

	// Pointer receiver method works on value — Go automatically takes address (&v).Scale(2)
	v.Scale(2)

	// Function requires explicit pointer
	ScaleFunc(&v, 10)

	// Even though Abs doesn't modify v, we still use a pointer receiver
	// This avoids copying the struct and ensures receiver consistency
	fmt.Println("v.Abs():", v.Abs())
	fmt.Println("AbsFunc(v):", AbsFunc(v))

	// ----- Pointer -----
	p := &Vertex{4, 3}

	// Method works on pointer
	p.Scale(3)

	// Function requires pointer
	ScaleFunc(p, 8)

	// Value receiver method works on pointer — Go automatically dereferences to (*p).Abs()
	fmt.Println("p.Abs():", p.Abs())
	fmt.Println("AbsFunc(*p):", AbsFunc(*p))

	// Final output of Vertex values
	fmt.Println("Final values:")
	fmt.Println("v =", v)
	fmt.Println("*p =", *p)
}
