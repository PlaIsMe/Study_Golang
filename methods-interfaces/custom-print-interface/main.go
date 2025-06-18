package main

import "fmt"

// Person defines a struct with Name and Age fields
type Person struct {
	Name string
	Age  int
}

// String implements the fmt.Stringer interface.
// It controls how Person is printed via fmt package functions.
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}

	// Because Person has a String() method, this prints customized output
	fmt.Println(a, z)
}
