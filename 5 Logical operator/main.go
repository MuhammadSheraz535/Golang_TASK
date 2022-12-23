package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	var x, y, z = 15, 25, 19

	// Arithmetic operators

	fmt.Println(x == y)
	fmt.Println(x != y)
	fmt.Println(x < y)
	fmt.Println(x <= y)
	fmt.Println(x > y)
	fmt.Println(x >= y)

	// Logical operators

	fmt.Println(x < y && x > z)
	fmt.Println(x < y || x > z)
	fmt.Println(!(x == y && x > z))

	now := time.Now()
	fmt.Println(now.Local().Year())

	var age string
	fmt.Println("Enter name")
	reader := bufio.NewReader(os.Stdin)
	age, err := reader.ReadString('\n')
	if err != nil {

		log.Fatal(err)

	}
	fmt.Println("Welcome ", age)

}
