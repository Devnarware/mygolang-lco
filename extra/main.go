package main

import "fmt"





func main()  {

	fmt.Println(max(3,4))
	
}


func max[T int | float64](a T, b T) T {
	if a>b{
		return a 
	}
	return b
}