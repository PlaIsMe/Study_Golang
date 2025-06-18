package main

import "fmt"

func describe(i interface{}) {
	fmt.Printf("describe: value = %v, type = %T\n", i, i)
}

func main() {
	var i interface{} // The empty interface can hold values of any type

	describe(i) // nil

	i = 42
	describe(i) // int

	i = "hello"
	describe(i) // string

	i = 3.14
	describe(i) // float64

	i = []string{"a", "b"}
	describe(i)
}
