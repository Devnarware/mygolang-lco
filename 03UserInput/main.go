package main

import (
	"bufio"
	"fmt"
	"os"
)



func main()  {
	fmt.Println(" ") ;
	fmt.Print("enter something: ") ;
	read := bufio.NewReader(os.Stdin) ;


	// Comma, ok || comma error

	input , _ := read.ReadString('\n') ;

	fmt.Println("entered value is: " + input)
}