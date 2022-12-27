package main

import (
	"fmt"
	"strings"
)

func show(array [7]string) {
	for k, v := range array {
		fmt.Println(k+1, v)

	}
}
func main() {
	// 16.1. Declaring arrays and accessing their elements
	array := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	for key, value := range array {
		fmt.Println(key+1, strings.ToUpper(value))
	}
	// Arrays are copied
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupitefmt",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	planetsMarkII := planets
	planets[2] = "whoops"
	fmt.Println(planets)
	fmt.Println(planetsMarkII)

	days := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	// passing an array to the function
	show(days)
	fmt.Printf("\n2D Array\n")

	//  Arrays of arrays
	board := [2][2]string{
		{"Afraz", "Shery"},
		{"Alaska", "Ali"}}

	for i := 0; i < len(board); i++ {

		for j := 0; j < len(board); j++ {
			fmt.Println(board[i][j])

		}

	}

}
