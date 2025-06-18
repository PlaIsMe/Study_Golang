package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// strings.NewReader returns a value that implements io.Reader
	r := strings.NewReader("Hello, Reader!")

	// Create a byte slice buffer of size 8
	b := make([]byte, 8)

	for {
		// Read fills b with data and returns number of bytes read and error
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n]) // print only the read portion

		if err == io.EOF {
			// End of stream
			break
		}
	}
}
