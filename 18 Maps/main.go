package main

import "fmt"

func main() {

	// 19.1. Declaring a map

	x := map[int]string{1: "Monday", 2: "Tuesday", 3: "Wednesday", 4: "Thursday"}
	// 19.3. Preallocating maps with make
	y := make(map[int]string)
	y[12] = "A"
	y[21] = "B"
	y[3] = "B"
	y[4] = "D"

	for k, v := range x {
		fmt.Println(k, v)
	}
	for k, v := range y {
		fmt.Println(k, v)
	}
	// 19.4 Using maps to count things
	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}
	// 19.5 Grouping data with maps and slices
	frequency := make(map[float64]int)
	for _, t := range temperatures {
		frequency[t]++
	}
	for t, num := range frequency {
		fmt.Printf("%+.2f occurs %d times\n", t, num)
	}
}

