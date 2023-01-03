package main

import "fmt"

type shape interface {
	area() float32
	perimeter() float32
}

type rectangle struct {
	length float32
	width  float32
}

type triangle struct {
	base   float32
	height float32
}

func (r rectangle) area() float32 {
	return r.length * r.width
}
func (r rectangle) perimeter() float32 {
	return 2 * r.length * r.width
}

func (t triangle) area() float32 {

	return 0.5 * t.base * t.height
}
func (t triangle) perimeter() float32 {

	return 2 * t.base * t.height
}
func calculate(s shape) float32 {
	return s.area()
}
func calculates(s shape) float32 {
	return s.perimeter()
}
func main() {
	rec := rectangle{3, 3}
	tria := triangle{4, 3}

	rect_area := calculate(rec)
	tri_area := calculate(tria)
	rect_per := calculates(rec)
	tri_per := calculates(tria)
	fmt.Println("Triangle area ", tri_area)
	fmt.Println("Rectangle area ", rect_area)
	fmt.Printf("Perimeter of rectangle : %v \n",rect_per)
	fmt.Printf("Perimeter of Triangle : %v \n",tri_per)

}
