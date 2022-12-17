package main

import (
	"fmt"
	"time"
)

func main() {

	//7.1. Declaring integer variables
	intVar := 9
	fmt.Printf("Type of variable is %T and value is %d", intVar, intVar)

	//7.2. The uint8 type for 8-bit colors

	var red, green, blue uint8 = 0, 141, 213
	fmt.Println(red, green, blue)

	a := "text"
	fmt.Printf("Type %T for %[1]s\n", a)
	b := 42
	fmt.Printf("Type %T for %[1]v\n", b)
	c := 3.14
	fmt.Printf("Type %T for %[1]v\n", c)
	d := true
	fmt.Printf("Type %T for %[1]v\n", d)

	// 7.3. Integers wrap around
	var yellow uint8 = 255
	yellow++
	fmt.Println(yellow)
	var number int8 = 127
	number++
	fmt.Println(number)

	future := time.Unix(12622780800, 0)
	fmt.Println(future)

}
