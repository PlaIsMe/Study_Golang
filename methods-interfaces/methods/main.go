package main

import (
	"fmt"
	"math"
)

// A method is like a function but with limited scope, making it clearer and more efficient

// ---- Struct with Method Receiver ----
type Vertex struct {
	X, Y float64
}

// Value receivers
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Pointer receivers
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// ---- Regular Function (No Method) ----
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Pointer receivers also work for function
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// ---- Method on Custom Type (not a struct) ----
type MyFloat float64

// Method for MyFloat type
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs()) // method call
	fmt.Println(Abs(v))  // function call

	// Pointer receivers will change the original value of v.
	// Value receivers will not, because v inside a value receiver method is a copy.
	// v inside a pointer receiver method is passed by reference (like a pointer).
	v.Scale(10)
	fmt.Println(v.Abs())

	Scale(&v, 10)
	fmt.Println(Abs(v))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs()) // method call on custom type
}
