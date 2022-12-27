package main

import (
	"fmt"
	"sort"
	"strings"
)

// 17.4. Slices with methods
func hyperspace(worlds []string) {
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}

// Declaring Varadic function
func terraform(prefix string, worlds ...string) []string {
	newWorlds := make([]string, len(worlds))
	for i := range worlds {
		newWorlds[i] = prefix + " " + worlds[i]
	}
	return newWorlds
}

func main() {
	// 17.1. Slicing an array
	planets := []string{" Venus ", "Earth ", " Mars"}
	//  18.1. The append function
	planets = append(planets, "Jupyters")
	hyperspace(planets)
	fmt.Println(strings.Join(planets, ""))
	// 18.2. Length and capacity
	fmt.Println("Length of an array =", len(planets))
	fmt.Println("Capacity of an array =", cap(planets))

	planetes := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}

	sort.StringSlice(planetes).Sort()
	fmt.Println(planetes)

	// Make a slice: slice-make.go
	dwarfs := make([]string, 0, 10)
	dwarfs = append(dwarfs, "Ceres", "Pluto", "Haumea", "Makemake", "Eris")
	fmt.Println(dwarfs)

	arr := []string{"Venus", "Mars", "Jupiter"}
	newPlanets := terraform("Hello", arr...)
	fmt.Println(strings.Join(newPlanets, "\n"))
}
