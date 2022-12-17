package main

import (
	"fmt"
	"strconv"
)

func main() {

	var floatValue float32 = 5.45

	// type conversion from float to int
	var intValue int = int(floatValue)

	fmt.Printf("Float Value is %g\n", floatValue)
	fmt.Printf("Integer Value is %d\n\n", intValue)

	fmt.Println("***************")

	var intValues int = 2

	// type conversion from int to float
	var floatValues float32 = float32(intValues)

	fmt.Printf("Integer Value is %d\n", intValues)
	fmt.Printf("Float Value is %f", floatValues)

	// Type conversion from string to int

	str := "200"

	intvar, _ := strconv.ParseInt(str, 0, 64)
	fmt.Printf("\nType of var is %T and value is %[1]v\n", intvar)

	// Type conversion from string to float

	str1 := "3.14"
	floatVar, _ := strconv.ParseFloat(str1, 64)
	fmt.Printf("\nType of var is %T and value is %[1]v\n", floatVar)

	// Type conversion from string to bool
	s1 := "true"
	b1, _ := strconv.ParseBool(s1)
	fmt.Printf("%T, %v\n", b1, b1)
}
