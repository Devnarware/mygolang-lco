package main

import "fmt"


type Shape interface{
	area() float64
}

type Rectangle struct {
	width  int
	height int
}
type Circle struct {
	radius int
}

func (r Rectangle) area() float64 {
	return float64(r.height * r.width)
}
func (c Circle) area() float64 {
	return 0.5*(3.14 * float64(c.radius) * float64(c.radius))
}

func (r Rectangle) perim() int {
	return 2*(r.height + r.width)
}

func printArea(s Shape){
	fmt.Printf("the area is %.2f \n", s.area())
}

func main() {

	var r Rectangle
	r.height = 10
	r.width = 20

	var c Circle 
	c.radius = 2 

	printArea(r)
	printArea(c)

}
