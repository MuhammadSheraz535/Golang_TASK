// Program to assign memory address to pointer

package main

import "fmt"

func reclassify(planets *[]string) {
	*planets = (*planets)[0:9]
}

func main() {

	var name = "John"
	var ptr *string

	//   26.1. The ampersand and the asterisk
	ptr = &name

	fmt.Println("Value of variable is", *ptr)
	fmt.Println("Address of the variable", ptr)

	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
		"Pluto",
	}
	reclassify(&planets)
	fmt.Println(&planets)

}
