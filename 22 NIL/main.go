package main

import (
	"fmt"
	"sort"
)

func main() {

	var soup []string
	var v interface{}
	var p *int
	v = p
	fmt.Printf("%T %v %v\n", v, v, v == nil)

	// slice of nil
	fmt.Println(soup == nil)
	soup = append(soup, "onion", "carrot", "celery")
	fmt.Println(soup)
	food := []string{"onion", "Apple", "carrot", "celery", "pae"}
	res := sort.StringSlice(food).Search("pae")
	if res == 0 {

		fmt.Println("NOT found")
	} else {
		fmt.Print(res)
	}
}
