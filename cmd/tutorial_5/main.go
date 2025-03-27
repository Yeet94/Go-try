package main

import "fmt"

type gasEngine struct {
	mpg     uint16
	gallons uint16
}

func (e gasEngine) milesLeft() uint16 {
	return e.gallons * e.mpg
}

func main() {
	var myCar gasEngine
	myCar.mpg = 30
	myCar.gallons = 10
	fmt.Println(myCar)
	fmt.Printf("Total miles left in tank: %v", myCar.milesLeft())
}
