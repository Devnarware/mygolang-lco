package main

import "fmt"



type car struct  {
	brand string ;
	model string ;
	price int ;

	fuel fuel ;

}

type fuel struct{
	fuelType string ;
}
func main()  {

	myCar := car{} ;

	myCar.brand = "tesla"
	myCar.model = "modelX" 
	myCar.price = 1000000 ;
	myCar.fuel.fuelType = "battery" ;

	

	printDetails(myCar)

}

func printDetails(car car)  {
	fmt.Println("")
	fmt.Println(car.brand)
	fmt.Println(car.model)
	fmt.Println(car.price)
	fmt.Println(car.fuel)
}

