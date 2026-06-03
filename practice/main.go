package main

import "fmt"

type Rectangle struct{
	width int
	height int
}

func (r Rectangle) area() int{
	return r.height * r.width
} 
func (r Rectangle) perim() int{
	return 2*(r.height + r.width)
} 

func main() {

	var a Rectangle 
	a.height = 10
	a.width = 20

	fmt.Println(a.area())
	fmt.Println(a.perim())

}


