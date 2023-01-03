package main

import "fmt"

func main() {

	age := make(chan int)
	name := make(chan string)
	go sendData(name, age)
	fmt.Printf("Name: %v\n", <-name)
	fmt.Printf("Name: %v\n", <-age)
	
}

func sendData(name chan string, age chan int) {

	name <-"Muhammad Sheraz"
	age <-19 
}