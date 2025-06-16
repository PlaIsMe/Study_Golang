package main

import "fmt"

// --- Define a custom type for map values ---
type Vertex struct {
	Lat, Long float64
}

func main() {
	// --- Create & use a map with make ---
	m1 := make(map[string]Vertex) // initialized map (empty)
	m1["Bell Labs"] = Vertex{40.68433, -74.39967}
	fmt.Println("m1:", m1["Bell Labs"])

	// --- Map Literal with explicit type ---
	m2 := map[string]Vertex{
		"Bell Labs": Vertex{40.68433, -74.39967},
		"Google":    Vertex{37.42202, -122.08408},
	}
	fmt.Println("m2:", m2)

	// --- Map Literal with shorthand (omit Vertex) ---
	m3 := map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println("m3:", m3)

	// --- Mutating a Map ---
	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println("Set:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("Update:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("Deleted:", m["Answer"]) // 0 (zero value)

	v, ok := m["Answer"] // check existence
	fmt.Println("Lookup:", v, "Present?", ok)
}
