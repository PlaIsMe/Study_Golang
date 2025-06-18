package main

import (
	"fmt"
	"math"
)

// Abser is an interface with one method signature
type Abser interface {
	Abs() float64
}

// MyFloat is a custom float type
type MyFloat float64

// MyFloat implements Abs with a value receiver
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return -float64(f)
	}
	return float64(f)
}

// Vertex is a struct with X, Y fields
type Vertex struct {
	X, Y float64
}

// *Vertex implements Abs with a pointer receiver
func (v *Vertex) Abs() float64 {
	if v == nil {
		fmt.Println("Abs called on nil *Vertex")
		return 0
	}
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// describe prints the interface's value and type
func describe(i Abser) {
	fmt.Printf("describe: value = %v, type = %T\n", i, i)
}

func main() {
	var a Abser

	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}
	var nil *Vertex

	a = f // ✅ MyFloat implements Abser
	describe(a)

	a = &v // ✅ *Vertex implements Abser
	describe(a)

	// a = v // ❌ Vertex does NOT implement Abser
	// Because Abs is defined on *Vertex only

	describe(nil) // <nil>

	fmt.Println("Abs:", a.Abs())
}

/*
==========================
🧠 INTERFACE NOTES
==========================

👉 What is an interface?
An interface is a type defined by a set of method signatures.
A value of interface type can hold any value that implements those methods.

✅ Interfaces are implemented implicitly.
A type implements an interface by implementing its methods.
There is no explicit declaration of intent — no "implements" keyword.

This means:
- You don’t need to say “type T implements interface I”
- Go checks if the method set matches automatically
- This decouples interface definition from its implementation,
  even across different packages.

✅ Under the hood:
An interface value is a pair: (value, type)

    (value, type)

Calling a method on an interface value dispatches to the method of the **underlying concrete type**.

📌 Example:
The function:
    func describe(i Abser) {
        fmt.Printf("(%v, %T)\n", i, i)
    }

Will show both the value and the dynamic type stored in the interface.
*/
