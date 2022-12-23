package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// 9.1. Declaring string variables
	var peace string = "peace"
	fmt.Println(peace)

	fmt.Println(`
 Hello, world	
 peace be upon you
 upon you be peace`)
	//  9.2  Characters, code points, runes, and bytes

	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33
	fmt.Printf("%c%c%c%c\n", pi, alpha, omega, bang)

	// 9.3 Pulling the strings

	var star byte = '*'
	fmt.Printf("%c %[1]v\n", star)
	smile := '😂'
	fmt.Printf("%c %[1]v\n", smile)
	acute := 'é'
	fmt.Printf("%c %[1]v\n", acute)

	// 9.4 Manipulating characters with Caesar cipher

	message := "shalom"
	for i := 0; i < 6; i++ {
		c := message[i]
		fmt.Printf("%c\n", c)
	}

	// 9.5. Decoding strings into runes

	question := "¿Cómo estás?"
	fmt.Println(len(question), "bytes")
	fmt.Println(utf8.RuneCountInString(question), "runes")
	c, size := utf8.DecodeRuneInString(question)
	fmt.Printf("First rune: %c %v bytes", c, size)

}
