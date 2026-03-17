package main

import (
	"bufio"
	"fmt"
	"os"
)


func main()  {
	

	fmt.Println("") ;
	fmt.Println("welcome to our app....") ;
	fmt.Print("rate the pizza from 1 to 5:")

	reader := bufio.NewReader(os.Stdin) ;

	intput, _ := reader.ReadString('\n') ;

	fmt.Println("you rated " + intput)
}