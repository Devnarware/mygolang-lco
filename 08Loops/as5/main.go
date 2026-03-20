package main

import (
	"fmt"
)

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i ==  0{
			return false
		}
	}
	return true 
}

func printPrimes(max int) {
	fmt.Println(2)
	for i := 3; i <= max; i++ {
		if isPrime(i){
			fmt.Println(i)
		}
	}
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}
