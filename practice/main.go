package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["Dev"] = 20
	m["Ved"] = 18
	m["Khed"] = 10

	for key, val := range m {
		if val >= 18 {
			fmt.Println(key)
		}
	}

}


