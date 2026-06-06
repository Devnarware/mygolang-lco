package main

import "fmt"

func main() {

	arr := []int {1,2,3,4,5,6}


	clearSlice(&arr)
	fmt.Println(arr)

}

func clearSlice(arr *[]int){
	*arr = []int {}
}


