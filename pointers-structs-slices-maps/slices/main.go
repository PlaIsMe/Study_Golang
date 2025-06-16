package main

import (
	"fmt"
	"strings"
)

func main() {
	// --- Slice from Array ---
	primes := [6]int{2, 3, 5, 7, 11, 13}
	s := primes[1:4] // s = [3, 5, 7]; shares data with primes
	fmt.Println("sliced primes:", s)

	// --- Shared Slice Memory ---
	names := [4]string{"John", "Paul", "George", "Ringo"}
	b := names[1:3]
	b[0] = "XXX" // names[1] modified → affects both a and b
	fmt.Println("shared names:", names)

	// --- Slice Struct (on stack) ---
	// A slice is: ptr to array (in heap), len, cap
	// Copying a slice copies this struct, not the data
	// Only appending beyond cap allocates a new array

	// --- Slice Literal ---
	nums := []int{2, 3, 5, 7} // underlying array in heap
	fmt.Println("slice literal:", nums)

	// --- Slice Defaults ---
	part := nums[1:3]
	fmt.Println("nums[1:3]:", part)
	fmt.Println("part[:1]:", part[:1], "part[1:]:", part[1:])

	// --- Length & Capacity ---
	s2 := []int{1, 2, 3, 4, 5}
	printSlice("s2", s2)
	s2 = s2[:0] // len = 0, cap = 5
	s2 = s2[:2] // len = 2
	s2 = s2[1:] // drop first
	printSlice("s2 trimmed", s2)

	// --- Nil Slice ---
	var nilSlice []int
	fmt.Println("nil slice:", nilSlice, len(nilSlice), cap(nilSlice))
	if nilSlice == nil {
		fmt.Println("nilSlice is nil")
	}

	// --- make() allocates slice ---
	a2 := make([]int, 5)    // len = 5, cap = 5
	b2 := make([]int, 0, 5) // len = 0, cap = 5
	c2 := b2[:2]            // len = 2, uses existing capacity
	d2 := c2[1:3]           // slice of slice
	printSlice("a2", a2)
	printSlice("b2", b2)
	printSlice("c2", c2)
	printSlice("d2", d2)

	// --- 2D Slice Example ---
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	board[0][0], board[1][2], board[2][2] = "X", "X", "O"
	for _, row := range board {
		fmt.Println(strings.Join(row, " "))
	}

	// --- Append Example ---
	var s3 []int
	s3 = append(s3, 1, 2, 3) // allocates if nil or cap exceeded
	printSlice("appended s3", s3)

	// --- Range Over Slice ---
	pow := []int{1, 2, 4, 8, 16}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// --- Range Tricks ---
	for i := range pow { // only index
		pow[i] = 1 << uint(i)
	}
	for _, v := range pow { // only value
		fmt.Println(v)
	}
}

func printSlice(name string, s []int) {
	fmt.Printf("%s: len=%d cap=%d %v\n", name, len(s), cap(s), s)
}
