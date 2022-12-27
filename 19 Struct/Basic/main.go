// Declaring a structure

package main

import (
	"fmt"
)

type curiosity struct {
	lat  float64
	long float64
}

// 23.1. Composing structures
type location struct {
	lat, long float64
	curiosity curiosity
}

type rectangle struct {
	length int
	width  int
}

// 22.1. Attaching methods to structures
func (r rectangle) area() int {

	return r.length * r.width
}

func main() {

	// 21.2. Reusing structures with types
	var c curiosity

	c.lat = -4.5895
	c.long = 137.4417
	fmt.Println(c.lat, c.long)
	c.lat = 78.8

	fmt.Println(c)
	cap := location{lat: 78.78, long: 90.89}
	// 21.4. Structures are copied
	cap1 := cap
	cap1.lat = 79.79
	cap1.long = 90.78
	cap.curiosity.long = 21.8
	cap.curiosity.lat = 12.1
	fmt.Println(cap, cap1)
	// 22.2. Constructor functions
	rect1 := rectangle{
		length: 5,
		width:  7,
	}
	rect2 := rectangle{
		length: 7,
		width:  7,
	}
	fmt.Println("Area of rectangle 1 = ", rect1.area())
	fmt.Println("Area of rectangle 2 = ", rect2.area())

}
