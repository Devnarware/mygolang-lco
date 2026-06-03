package main

import "fmt"

func main() {

	var a, b int
	fmt.Print("Enter the value of a and b: ")
	fmt.Scan(&a, &b)

	ans := add(a, b)
	fmt.Println(ans)

}

func add(a int, b int) int {
	return a + b
}
