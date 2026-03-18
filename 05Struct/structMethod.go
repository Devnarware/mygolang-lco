package main

import "fmt"

type rect struct{
	width int 
	height int 
}

func (r rect) area() int {

	area := r.height * r.width
	return area
}


func main()  {

	var r rect 

	r.height = 2
	r.width = 2

	fmt.Println(" ")
	fmt.Println("the area of the rectangle is" , r.area())

}





