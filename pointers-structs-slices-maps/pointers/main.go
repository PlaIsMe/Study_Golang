package main

import "fmt"

func main() {
	i, j := 42, 2701

	var p *int      // Declare a pointer to int (default value is nil)
	p = &i          // p holds the address of i
	fmt.Println(*p) // Dereference p: read the value at the address (prints 42)
	*p = 21         // Dereference p: update the value at the address
	fmt.Println(i)  // i is now 21

	p = &j         // p now points to j
	*p = *p / 37   // Modify j via the pointer (2701 / 37 = 73)
	fmt.Println(j) // j is now 73

	// Note: Go does NOT allow pointer arithmetic like p++, p--, or p + 1 (unlike C/C++)
}
