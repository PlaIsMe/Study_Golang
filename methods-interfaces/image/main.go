package main

import (
	"fmt"
	"image"
)

func main() {
	// Create a new RGBA image with a 100x100 bounding box
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))

	// Print the bounds of the image
	fmt.Println(m.Bounds()) // Output: (0,0)-(100,100)

	// Print the RGBA values of the pixel at (0, 0)
	fmt.Println(m.At(0, 0).RGBA()) // Output: (0, 0, 0, 0)
}
