package main

import "fmt"

func main() {
	func() {
		fmt.Printf("Anonymous function\nFunction with no name")
	}()
}
