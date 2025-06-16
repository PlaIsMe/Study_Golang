package main

import "fmt"

// func <function_name>(<parameters>) <return type>
// The parameter name comes before the type, which is the reverse of C-style syntax.
// A function can be shortened when parameters share the same type, e.g., func add(x, y int)
func add(x int, y int) int {
	return x + y
}

// A function can return multiple results.
func swap(x, y string) (string, string) {
	return y, x
}

// You can name the return values. In that case, you can omit them in the return statement.
// Go will understand and return the named values.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(42, 13))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))
}
