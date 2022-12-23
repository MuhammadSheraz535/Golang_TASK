package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Bring your own type
type kelvin float64

func measureTemp(samples int, temp func() kelvin) {
	for i := 0; i < samples; i++ {
		k := temp()
		fmt.Println(k, "K")
		time.Sleep(time.Millisecond * 100)
	}

}

func RandTempGeneartor() kelvin {

	rand.Seed(time.Now().UnixNano())
	return kelvin(rand.Intn(150) + 150)

}

func main() {
	// Passing the function as argument
	measureTemp(5, RandTempGeneartor)
	//  Anonymous function
	z := func() {
		fmt.Printf("Random function\n")
	}
	//Assigning functions to variables
	y := z
	y()

}
