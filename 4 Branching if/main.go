// Program to relate two integers using =, > or < symbol

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter marks :")

	marks := bufio.NewReader(os.Stdin)
	input, err := marks.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	grade , err := strconv.ParseFloat(input,64)
	if err != nil {
		log.Fatal(err)
	}
	var  status string
	if grade >40{
		status = "Passed"
	}else{
		status = "Failed"
	}
	fmt.Printf("You have %s with %v  marks",status,grade)

}
