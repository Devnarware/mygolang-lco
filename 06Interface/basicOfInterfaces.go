package main

import (
	"math"
	"fmt"
)


type shape interface{
	area() float64
	permeter() float64
}


// FOR THE RECTANGLE 

type rectanle struct{
	length, width int
}

func (r rectanle) area() float64 {
	area := float64(r.length * r.width)
	return area 
}



// FOOR THE CIRCLE	

type circle struct{
	radius int 
}

func (c circle) perimeter() float64 {
	peri := 2 * math.Pi * float64(c.radius)
	return peri 
}



func main()  {

	var rect rectanle

	rect.width = 3
	rect.length = 7

	fmt.Println("")
	fmt.Println(rect.area())

}
