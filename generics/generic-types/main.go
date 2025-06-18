package main

import "fmt"

// SumInts adds together the values of a map with int64 values.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of a map with float64 values.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// Number is a type constraint that matches int64 or float64.
type Number interface {
	int64 | float64
}

// SumNumbers is a generic function that sums map values of type int64 or float64.
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	ints := map[string]int64{"a": 1, "b": 2, "c": 3}
	floats := map[string]float64{"a": 1.5, "b": 2.5, "c": 3.5}

	fmt.Println("SumInts:", SumInts(ints))                  // 6
	fmt.Println("SumFloats:", SumFloats(floats))            // 7.5
	fmt.Println("SumNumbers[int64]:", SumNumbers(ints))     // 6
	fmt.Println("SumNumbers[float64]:", SumNumbers(floats)) // 7.5
}
