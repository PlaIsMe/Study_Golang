package main

import "fmt"

// Type assertion demo: access the concrete value inside an interface
func typeAssertionDemo() {
	fmt.Println("== Type Assertion Demo ==")

	var i interface{} = "hello"

	// Unsafe assertion (panics if wrong)
	s := i.(string)
	fmt.Println("asserted string:", s)

	// Safe assertion with check
	s, ok := i.(string)
	fmt.Println("string:", s, "| success:", ok)

	f, ok := i.(float64) // i is not float64
	fmt.Println("float64:", f, "| success:", ok)

	// Uncomment below to trigger panic
	// f = i.(float64)
	// fmt.Println("unsafe float64:", f)
}

// Type switch demo: perform actions depending on actual type
func typeSwitchDemo(i interface{}) {
	fmt.Println("== Type Switch Demo ==")

	switch v := i.(type) {
	case int:
		fmt.Printf("int: Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("string: %q is %d bytes long\n", v, len(v))
	case float64:
		fmt.Printf("float64: %.2f squared is %.2f\n", v, v*v)
	default:
		fmt.Printf("unknown type: %T\n", v)
	}
}

func main() {
	typeAssertionDemo()

	fmt.Println()

	typeSwitchDemo(21)
	typeSwitchDemo("hello")
	typeSwitchDemo(3.14)
	typeSwitchDemo(true)
}
