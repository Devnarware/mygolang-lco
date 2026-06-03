package main

import "fmt"

func main() {

	ans := add(1, 2)
	fmt.Println(ans)
}

func add(a int, b int) int {
	return a + b
}
