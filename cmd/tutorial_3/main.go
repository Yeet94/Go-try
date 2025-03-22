package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "Hello World!"
	printMe(printValue)

	var intNum int = 42
	var intNum2 int = 8
	var result, remainder, err = intDevision(intNum, intNum2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("The result of the integer division is %v with remainder %v.\n", result, remainder)
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDevision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("division by zero is not allowed")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
