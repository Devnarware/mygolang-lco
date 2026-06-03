package main

import "fmt"

func main() {

	ans := add(1, 2)
	fmt.Println(ans)
	fmt.Println("hello")
}

func add(a int, b int) int {
	return a + b
}
