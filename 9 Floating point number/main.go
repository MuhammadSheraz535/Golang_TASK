package main

import (
	"fmt"
	"math"
)

func main() {

	// 6.1. Declaring floating-point variables
	var pi64 = math.Pi
	fmt.Println(pi64)

	// 6.2. Displaying floating-point types
	fmt.Printf("Type of variablr is %T\n", pi64)

	//6.3. Floating-point accuracy

	third := 1.0 / 3.0
	fmt.Println(third + third + third)

	// 6.4. Comparing floating-point numbers

	piggyBank := 0.1
	piggyBank += 0.2
	fmt.Println(piggyBank == 0.3)

	
	fmt.Println(math.Abs(piggyBank-0.3) < 0.0001)



}
