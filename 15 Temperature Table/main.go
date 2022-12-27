package main

import "fmt"

func F_TO_C() {

	fmt.Print("=============================== \n")
	fmt.Printf("|  F° \t\t|| \tC°    |\n")
	fmt.Print("=============================== \n")
	for i := -40.0; i <= 100.0; i = i + 5.0 {
		c := ((i - 32) * 5 / 9)
		fmt.Printf("|%-5.1fF°\t\t%-5.1fC°|\n", i, c)

	}

	fmt.Print("===============================")
}
func C_TO_F() {

	fmt.Print("\n=============================== \n")
	fmt.Printf("|  C° \t\t|| \tF°    |\n")
	fmt.Print("=============================== \n")
	for i := -40.0; i <= 100.0; i = i + 5.0 {
		c := ((i*9)/5 + 32)
		fmt.Printf("|%-5.1fF°\t\t%-5.1fC°|\n", i, c)

	}

	fmt.Print("===============================")
}

func main() {
	F_TO_C()
	C_TO_F()

}
